package liquidplanner



/*******************************************************************************
 * Model
 */

type Model struct {
  Id int32          `json:"id"`
}



/*******************************************************************************
 * Workspace
 */

type Workspace struct {
  Model

  Name string
}



/*******************************************************************************
 * Task
 */

type Task struct {
  Model

  Name string       `json:"name",omitempty`
  Owner_id int32    `json:"owner_id"`
}



/*******************************************************************************
 * Project
 */

type Project struct {
  Model

  Name string
}



/*******************************************************************************
 * Person
 */

type Person struct {
  Model

  First_name string
  Last_name string

  Workspaces []Workspace
}

func ( self *Person ) FullName ( ) ( string ) {
  return self.First_name + " " + self.Last_name
}
