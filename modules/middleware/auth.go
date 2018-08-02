package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"eth-agent/config"
)

var (
	httpClient = &http.Client{}
)

// ProxyResp - The response format of eth proxy
type AgentResp struct {
	Msg string `json:"msg"`
}

// ProxyResp - The response format of eth proxy
type OtherResp struct {
	Msg string `json:"msg"`
}

type Res struct {
	Msg         string `json:msg`
	Status      bool
	User_id     string `json:user_id`
	Expire_time int64
}

// VerifyToken - check if the token is valid
// router.POST("/:token", redirect)
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		url := c.Request.URL.String()
		fmt.Println(url)
		if token == "" {
			startsWithAgent := strings.HasPrefix(url, "/agent")
			startsWithOther := strings.HasPrefix(url, "/other")
			if startsWithAgent {
				c.JSON(http.StatusUnauthorized, AgentResp{Msg: "Access token should not be empty."})
				c.Abort()
				return
			} else if startsWithOther {
				c.JSON(http.StatusUnauthorized, OtherResp{Msg: "Access token should not be empty."})
				c.Abort()
				return
			}
		}

		postData, err := json.Marshal(map[string]string{
			"token_id": token,
			"from":     "sso_scope",
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		req, err := http.NewRequest(http.MethodPost, config.ToeknVerifyURI, bytes.NewBuffer(postData))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		var arbitraryObj Res
		json.Unmarshal(body, &arbitraryObj)

		// The status code will be 400 when token verify failed
		if !arbitraryObj.Status {
			startsWithAgent := strings.HasPrefix(url, "/agent")
			startsWithOther := strings.HasPrefix(url, "/other")
			verifyMsg := arbitraryObj.Msg
			if startsWithAgent {
				c.JSON(http.StatusUnauthorized, AgentResp{Msg: verifyMsg})
				c.Abort()
				return
			} else if startsWithOther {
				c.JSON(http.StatusUnauthorized, OtherResp{Msg: verifyMsg})
				c.Abort()
				return
			}
		}
	}
}
