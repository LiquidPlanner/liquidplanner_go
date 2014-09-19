package liquidplanner

import (
  "strconv"
)

//------------------------------------------------------------------------------
//-- Account
//------------------------------------------------------------------------------

func GetAccount() ( result Person, res LpResponse ) {
  res = FetchJson( "GET", "account", nil, &result )

  return
}

//------------------------------------------------------------------------------
//-- Workspace
//------------------------------------------------------------------------------

func GetWorkspaces() ( result []Workspace ) {
  FetchJson( "GET", "workspaces", nil, &result )

  return
}

//------------------------------------------------------------------------------
//-- Project
//------------------------------------------------------------------------------

func GetProjects() ( result []Project ) {
  FetchJson( "GET", "workspaces/" + lp_space_id + "/projects", nil, &result)

  return
}

//------------------------------------------------------------------------------
//-- Task
//------------------------------------------------------------------------------

// get all
func GetTasks( ) ( result []Task ) {
  FetchJson( "GET", "workspaces/" + lp_space_id + "/tasks", nil, &result)

  return
}

// create
func CreateTask( task Task ) ( result Task ) {
  taskWrap := map[string]Task { "task": task, } 

  FetchJson( "POST", "workspaces/" + lp_space_id + "/tasks", taskWrap, &result)

  return
}

// update
func UpdateTask( task *Task ) {
  taskWrap := map[string]Task { "task": *task, } 

  FetchJson( "PUT", "workspaces/" + lp_space_id + 
                   "/tasks/" + strconv.FormatInt( int64(task.Id), 10 ), 
             taskWrap, task)

  return
}

// delete
func DeleteTask( task *Task ) {
  Fetch( "DELETE", "workspaces/" + lp_space_id + 
                   "/treeitems/" + strconv.FormatInt( int64(task.Id), 10 ), 
         nil )
}
