package handler

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"strconv"
	"../../customlogger"
	"../../database"
	api_middleware "../mymiddleware"
	red "../redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// formlogin
func FormSignIn(c echo.Context) error {
	cookie, err := c.Cookie(api_middleware.COOKIE_NAME)

	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			customlogger.SetUser("LOGIN")
			customlogger.Debug("Login Page")
			return c.Render(http.StatusOK, "form_login", nil)
		}

		log.Println("err cookie signin")
		log.Println(err)
		return err
	}

	if cookie.Value == "" {

		customlogger.SetUser("LOGIN")
		customlogger.Debug("Login Page")
		return c.Render(http.StatusOK, "form_login", nil)

	}

	return c.Redirect(301, "/lib/dashboard/")

}

//login proses
func AuthorizationSignIn(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	// gob.Register(UserModel{})

	formusername := c.FormValue("username")
	formpassword := c.FormValue("password")

	//start cek database
	var salt, password, username string
	sqlStatement := `SELECT salt, password, username FROM v_get_user_grup WHERE  username = ?`
	row := db.QueryRow(sqlStatement, formusername)
	err := row.Scan(&salt, &password, &username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=username_false")
		} else {
			panic(err)
			fmt.Println(err)

		}
	}
	//end check database

	// MD5 PASSWORD

	var str_password string = formpassword

	hasher_password := md5.New()
	hasher_password.Write([]byte(str_password))
	md5password := hex.EncodeToString(hasher_password.Sum(nil))

	// set password dan salt
	var salt_password string = salt + md5password

	hasher_salt_password := md5.New()
	hasher_salt_password.Write([]byte(salt_password))
	get_password := hex.EncodeToString(hasher_salt_password.Sum(nil))

	//validasi password
	if get_password != password && username != formusername {
		return c.String(http.StatusInternalServerError, "not same password")
	}
	//end validasi password


	var id_setting_user, id_setting_grup, full_name, name_grup, image, extension []byte
	sqlStatement = `SELECT id_setting_user, id_setting_grup, full_name, name_grup, image, extension FROM v_get_user_grup WHERE  username = ? AND password = ?`
	row = db.QueryRow(sqlStatement, username, get_password)
	err = row.Scan(&id_setting_user, &id_setting_grup, &full_name, &name_grup, &image, &extension)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=password_false")
		} else {
			panic(err)
			fmt.Println(err)

		}
	}

	//TODO create jwt token
	struct_claims := api_middleware.JwtClaims{
		string(id_setting_user),
		string(id_setting_grup),
		string(full_name),
		string(name_grup),
		string(image),
		string(extension),
		jwt.StandardClaims{
			Id:        "main_" + string(id_setting_user),
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token, err := api_middleware.CreateJwtToken(struct_claims)
	if err != nil {
		log.Println("error create jwt token", err)
		return err
	}
	//end create token JWT


	//start convert md5
	keyCookie := "martin" + string(id_setting_user)
	str := keyCookie
	hasher := md5.New()
	hasher.Write([]byte(str))
	conv_key_martin := hex.EncodeToString(hasher.Sum(nil))

	keyCookie = "redis" + string(id_setting_user)
	str = keyCookie
	hasher = md5.New()
	hasher.Write([]byte(str))
	conv_key_redis := hex.EncodeToString(hasher.Sum(nil))

	get_key_redis := conv_key_martin + conv_key_redis

	//set key redis + cookie
	getKeyRedis := string([]rune(get_key_redis)[8:28])
	fmt.Println(getKeyRedis)
	//end set key redis + cookie

	//start set cookie
	cookie := &http.Cookie{}

	cookie.Name = api_middleware.COOKIE_NAME
	cookie.Value = getKeyRedis
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
	//end set cookie

	client := red.Connection()

	err = client.Set(getKeyRedis, token, 0).Err()
	if err != nil {
		fmt.Println("err set redis")
		fmt.Println(err)
		return err
	}

	//start set expire redis
	expire := client.Expire(getKeyRedis, 1*time.Hour)
	fmt.Println(expire.Val())
	//end set expire redis

	return c.Redirect(http.StatusTemporaryRedirect, "/lib/sign/redirect/")
}

func RedirectSignIn(c echo.Context) error {
	return c.Redirect(301, "/lib/dashboard/")
}

//logout
func AuthorizationSignOut(c echo.Context) error {

	cookie, err := c.Cookie(api_middleware.COOKIE_NAME)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return c.Redirect(http.StatusTemporaryRedirect, "/logout/redirect/")
		}

		log.Println("err cookie singout")
		log.Println(err)
		return err
	}

	if cookie.Value != "" {

		//start connection redis
		client := red.Connection()
		//end connection to redis

		//start remove redis where key
		index, err := client.Del(cookie.Value).Result()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(" Index ke-", index)
		//end remove redis where  key

		//star set expire redis
		boolean := client.Expire(cookie.Value, 0*time.Second)
		fmt.Println("status:", boolean)
		//end set expire redis

		//start set expire cookie
		expCookie := &http.Cookie{
			Name:    api_middleware.COOKIE_NAME,
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),

			HttpOnly: true,
		}
		c.SetCookie(expCookie)
		//end set expire cookie
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/logout/redirect/")
}

func RedirectLogout(c echo.Context) error {
	t   			   := time.Now().UnixNano()
	current_time 	   := strconv.FormatInt(t, 10)
	return c.Redirect(http.StatusTemporaryRedirect, "/?"+current_time+"")
}

func POSTRedirectLogout(c echo.Context) error {
	return c.Render(http.StatusOK, "form_login", nil)
}

// session_expire
func AuthorizationSession(c echo.Context) error {

	cookie, err := c.Cookie(api_middleware.COOKIE_NAME)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return c.Redirect(http.StatusTemporaryRedirect, "/?session_expire=true")
		}
		log.Println("err cookie singout")
		log.Println(err)
		return err
	}

	if cookie.Value != "" {

		//start connection redis
		client := red.Connection()
		//end connection to redis

		//start remove redis where key
		index, err := client.Del(cookie.Value).Result()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(" Index ke-", index)
		//end remove redis where  key

		//star set expire redis
		boolean := client.Expire(cookie.Value, 0*time.Second)
		fmt.Println("status:", boolean)
		//end set expire redis

		//start set expire cookie
		expCookie := &http.Cookie{
			Name:    api_middleware.COOKIE_NAME,
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),

			HttpOnly: true,
		}
		c.SetCookie(expCookie)
		//end set expire cookie
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/?session_expire=true")
}

func RedirectSession(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/?session_expire=true")
}