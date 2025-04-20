package repo

// func ValidateCredentials(username string, password string) error {
// 	query := "SELECT username, password FROM users WHERE username=$1"
// 	row := db.DB.QueryRow(query, username)

// 	var retrivedPassword, retrivedUsername string
// 	err := row.Scan(&retrivedUsername, &retrivedPassword)

// 	fmt.Printf("retrived username: %s", retrivedUsername)
// 	fmt.Printf("retrived password: %s", retrivedPassword)
// 	fmt.Printf(" username: %s", username)
// 	if err != nil {
// 		return errors.New("credentials invalid")
// 	}

// 	if retrivedUsername == username {
// 		if retrivedPassword == password {
// 			return nil
// 		}
// 	} else {
// 		return errors.New("Invalid UserName")
// 	}
// 	return nil

// }
