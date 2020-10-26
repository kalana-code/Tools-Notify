package app

import (
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/notify/auth"
	"github.com/notify/db"
	"os"
	"time"
)


//Insert Application
func InsertApp(applicationName string) (*ApplicationInfoResponse,error) {
	applicationId := uuid.New()

	// generate application key
	appKey := &auth.ApplicationKey{
		ApplicationId : applicationId.String(),
		ApplicationName:applicationName,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*24*6*365).Unix(),
		},
	}


	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), appKey)
	// sign token
	tokenString, err := token.SignedString([]byte(os.Getenv("jwtSecret")))
	if(err!=nil){
		return nil,err
	}

	
	dbConnection, err := db.GetMySQLConnection()
	if(err !=nil){
		return nil,err
	}

	stmtIns, err := dbConnection.Prepare(
		`INSERT INTO app_keys( application_id, application_name, application_access_key) VALUES(?,?,?)`)

	if err != nil {
		return nil,err
	}
	_, err = stmtIns.Exec(applicationId,applicationName,tokenString)

	if err != nil {
		return nil,err
	}

	stmtIns.Close()

	return &ApplicationInfoResponse{
		ApplicationName: applicationName,
		ApplicationId: applicationId.String(),
		ApplicationAccessToken: tokenString,
	},nil
}

//Find application by access-key
func FindApp(accessToken string) (*ApplicationInfo,error) {
	applicationInfo := &ApplicationInfo{};
	dbConnection, err := db.GetMySQLConnection()
	if(err !=nil){
		return nil, err
	}
	stmtOut, err := dbConnection.Prepare("SELECT application_id, application_name FROM app_keys WHERE application_access_key	 = ?")
	if err != nil {
		return nil, err
	}


	err = stmtOut.QueryRow(accessToken).Scan(
		&applicationInfo.ApplicationId,
		&applicationInfo.ApplicationName,
	)
	if err != nil {
		return nil, err
	}
	return applicationInfo, nil
}

func GetApplicationLevelsDetails(accessToken string)  ([]ApplicationLevels, error)  {
	applicationLevels := []ApplicationLevels{}
	dbConnection, err := db.GetMySQLConnection()
	if(err !=nil){
		return nil, err
	}
	rows,err := dbConnection.Query(`SELECT levels.subscribe_level_name,levels.subscribe_level  FROM notifications_system.app_keys  as app inner join app_subscribe_level as levels on levels.application_id = app.application_id where app.application_access_key=?`,accessToken)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		applicationLevel := ApplicationLevels{}
		err = rows.Scan(&applicationLevel.LevelName,&applicationLevel.LevelId)
		if err != nil {
			return nil, err
		}
		applicationLevels = append(applicationLevels,applicationLevel)
	}
	return applicationLevels, nil
}

