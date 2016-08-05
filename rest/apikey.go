package rest

// // GetApikeys returns a list of all API keys under the account
// func (c APIClient) GetApikeys() ([]*ns1.Apikey, error) {
// 	path := "account/apikeys"
// 	req, err := c.NewRequest("GET", path, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	apikeys := []*ns1.Apikey{}
// 	_, err = c.Do(req, &apikeys)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return apikeys, nil
// }

// // GetApikey takes an ID and returns details, including permissions, for a single API key
// func (c APIClient) GetApikey(id string) (*ns1.Apikey, error) {
// 	path := fmt.Sprintf("account/apikeys/%s", id)
// 	req, err := c.NewRequest("GET", path, nil)

// 	var k ns1.Apikey
// 	_, err = c.Do(req, &k)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &k, nil
// }

// // // CreateApikey takes an *Apikey and creates a new API key
// // func (c APIClient) CreateApikey(k *Apikey) error {
// // 	path := "account/apikeys"
// // 	req, err := c.NewRequest("PUT", path, &k)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	// Update apikey fields with data from api(ensure consistent)
// // 	_, err = c.Do(req, &k)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // // DeleteApikey takes an ID and deletes and API key
// // func (c APIClient) DeleteApikey(id string) error {
// // 	path := fmt.Sprintf("account/apikeys/%s", id)
// // 	req, err := c.NewRequest("DELETE", path, nil)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	_, err = c.Do(req, nil)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }

// // // UpdateApikey takes an *Apikey and change name or access rights for an API key
// // func (c APIClient) UpdateApikey(k *Apikey) error {
// // 	path := fmt.Sprintf("account/apikeys/%s", k.Id)

// // 	req, err := c.NewRequest("POST", path, &k)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	// Update apikeys fields with data from api(ensure consistent)
// // 	_, err = c.Do(req, &k)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return nil
// // }
