package server

import (
    "context"
    "encoding/json"
    "net/http"
    "os"
	"time"
	"log"

    "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

// OAuth configuration
var googleOauthConfig = &oauth2.Config{
    ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
    ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
    RedirectURL:  "http://localhost:8080/auth/google/callback",
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:     google.Endpoint,
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func init() {
    log.Printf("JWT_SECRET: %s", os.Getenv("JWT_SECRET"))
}

func (s *Server) HandleGoogleLogin(c *gin.Context) {
    url := googleOauthConfig.AuthCodeURL("random_state_string", oauth2.AccessTypeOffline)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func (s *Server) HandleGoogleCallback(c *gin.Context) {
    code := c.Query("code")
    if code == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
        return
    }

    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
        return
    }

    // Fetch user information using the token
    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
        return
    }
    defer resp.Body.Close()

    // Parse user data
    var userData struct {
        Email string `json:"email"`
        // Add other fields as needed
    }
    if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
        return
    }

    // Create or update user in the database
    // TODO: Implement user creation or lookup

    // Generate JWT token
    jwtToken, err := generateJWT(userData.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Redirect to frontend with token
    c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/dashboard?token="+jwtToken)

}

func generateJWT(email string) (string, error) {
    claims := jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}