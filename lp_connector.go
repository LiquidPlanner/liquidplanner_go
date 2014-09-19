package liquidplanner

import (
  "bytes"
  "encoding/json"
  "errors"
  "io"
  "io/ioutil"
  "net/http"
  "strconv"
)

//------------------------------------------------------------------------------
//-- Structure to hold a response
//------------------------------------------------------------------------------

type LpResponse struct {
  Response *http.Response
  Body string
  Error error
}

//------------------------------------------------------------------------------
//-- Fetch a request from LiquidPlanner
//------------------------------------------------------------------------------

func Fetch( method, urlStr string, data interface{} ) ( LpResponse ) {
  var body io.Reader

  if data != nil {
    b, _ := json.Marshal( data )
    body = bytes.NewBuffer( b )
  } else {
    body = nil
  }

  req := request( method, urlStr, body )

  if data != nil {
    req.Header.Add( "Content-Type", "application/json" )
  }

  return ProcessRequest( req )
}

//------------------------------------------------------------------------------
//-- Fetch a request from LiquidPlanner and create a model from the response
//------------------------------------------------------------------------------

func FetchJson( method, urlStr string, 
                data interface{}, rs interface{} ) ( res LpResponse ) {
  res = Fetch( method, urlStr, data )

  json.Unmarshal( []byte(res.Body), &rs )
  
  return
}

//------------------------------------------------------------------------------
//-- Create and authenticated request
//------------------------------------------------------------------------------

func request( method, urlStr string, body io.Reader ) ( req *http.Request ) {
  fullUrl := "https://app.liquidplanner.com/api/" + urlStr

  req, err := http.NewRequest( method, fullUrl, body )
  if ( err != nil ) { return nil }
  
  return req
}

//------------------------------------------------------------------------------
//-- Handle the result
//------------------------------------------------------------------------------

func ProcessRequest( req *http.Request ) ( LpResponse ) {
  client := &http.Client { }

  req.SetBasicAuth( username, password )

  resp, err := client.Do( req )
  defer resp.Body.Close()

  var lp_response LpResponse

  lp_response.Response = resp

  if err != nil {
    lp_response.Error = err
  } else if resp.StatusCode >= 300 {
    lp_response.Error = errors.New( resp.Status )
  }
  
  if b, err := ioutil.ReadAll(resp.Body); err == nil {
    lp_response.Body = string(b)
  } 

  return lp_response
}

//------------------------------------------------------------------------------
//-- Username and password
//------------------------------------------------------------------------------

var username string
var password string

func Login( user string, pass string ) {
  username = user
  password = pass
}

//------------------------------------------------------------------------------
//-- Space Id accessors
//------------------------------------------------------------------------------

var lp_space_id string

func SetSpaceId( spaceId int32 ) {
  lp_space_id = strconv.FormatInt(int64(spaceId), 10)
}

func SetSpace( space Workspace ) {
  lp_space_id = strconv.FormatInt(int64(space.Id), 10)
}

//------------------------------------------------------------------------------

