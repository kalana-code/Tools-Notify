package subscriber

import (
	"database/sql"
	"errors"
	"github.com/notify/db"
	"log"
)


// subscribe application utils
//-------------------------------------------------------------------
func SubscribeApp(subReq SubscribeRequest ,appId string) (error) {
	//Level count is zero then return error
	if(subReq.LevelCount<1){
		log.Println("print Info")
		return  errors.New("Level Count cannot be zero or negative")
	}
	db, err := db.GetMySQLConnection()
	if err != nil {
		return  err
	}
	// begin a transaction
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	//create subscriber
	err = subscribe(tx, subReq.Email, subReq.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	//get subscriber id and map with application
	err = mapWithApplication(tx,subReq.Email,appId,subReq.LevelCount)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return  nil
}

func getAppContext(data interface{}) (*ApplicationContext, error) {
	ctx := data.(map[string]string)
	applicationContext := ApplicationContext{}
	appName, ok := ctx["ApplicationName"]
	if ok {
		applicationContext.ApplicationName = appName
	} else {
		return nil, errors.New("Cannot find application proper context")
	}

	appID, ok := ctx["ApplicationID"]
	if ok {
		applicationContext.ApplicationID = appID
	} else {
		return nil, errors.New("Cannot find application proper context")
	}
	return &applicationContext, nil
}

func subscribe(db *sql.Tx, email string, name string) (error){
	stmtIns, err := db.Prepare(
		`
			INSERT INTO subscribers ( subscriber_email,subscriber_name)
			VALUES (?,?)
			ON DUPLICATE KEY UPDATE subscriber_name = ?
		`)
	if err != nil {
		return err
	}
	defer  stmtIns.Close()
	_, err = stmtIns.Exec(
		email,
		name,
		name,
	)
	if err != nil {
		return  err
	}
	return nil
}

func mapWithApplication(db *sql.Tx, email string,appId string,level int) (error){
	var subscriberID int
	stmtOut, err := db.Prepare("SELECT subscriber_id  FROM subscribers where subscriber_email= ?")
	if err != nil {
		return err
	}
	defer  stmtOut.Close()

	err = stmtOut.QueryRow(email).Scan(
		&subscriberID,
	)
	if err != nil {
		return err
	}


	stmtIns, err := db.Prepare(
		`
			INSERT INTO subscriber_map ( subscriber_id, subscriber_app_id, subscriber_levels_count)
			VALUES (?,?,?)
			ON DUPLICATE KEY UPDATE subscriber_levels_count = ?
		`)
	if err != nil {
		return err
	}
	defer  stmtIns.Close()
	_, err = stmtIns.Exec(
		subscriberID,
		appId,
		level,
		level,
	)
	if err != nil {
		return  err
	}
	return nil
}

// end subscribe functions

// get subscriber list for given application and list
func SubscribersList(level int , appId string) (*[]Subscriber,error){
	//check valid level
	if(!isPowerOfTwo(level)){
		return  nil,errors.New("Invalid application level. Level always should be power of two,")
	}

	subscribersList := []Subscriber{}
	dbConnection, err := db.GetMySQLConnection()
	if(err !=nil){
		return nil,err
	}
	rows,err := dbConnection.Query(
		`SELECT
				su.subscriber_id  as id,
					su.subscriber_email as email,
					sm.subscriber_levels_count as count
				FROM subscriber_map as  sm
				inner join subscribers as su
				on sm.subscriber_id = su.subscriber_id
				where sm.subscriber_app_id = ? and su.subscriber_state='ACTIVE' and sm.subscriber_levels_count & ?
				`,appId,level,
	)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	for rows.Next() {
		subscriber := Subscriber{}
		err = rows.Scan(&subscriber.Id,&subscriber.Email,&subscriber.LevelCount)
		if err != nil {
			return nil,err
		}
		subscribersList = append(subscribersList,subscriber)
	}

	return &subscribersList,nil
}

func isPowerOfTwo(level int) bool{
	return (level !=0) && ((level &(level -1))==0)
}

