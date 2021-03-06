// Package fusiontables provides access to the Fusion Tables API.
//
// See https://developers.google.com/fusiontables
//
// Usage example:
//
//   import "google.golang.org/api/fusiontables/v1"
//   ...
//   fusiontablesService, err := fusiontables.New(oauthHttpClient)
package fusiontables

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "fusiontables:v1"
const apiName = "fusiontables"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/fusiontables/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage your Fusion Tables
	FusiontablesScope = "https://www.googleapis.com/auth/fusiontables"

	// View your Fusion Tables
	FusiontablesReadonlyScope = "https://www.googleapis.com/auth/fusiontables.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Column = NewColumnService(s)
	s.Query = NewQueryService(s)
	s.Style = NewStyleService(s)
	s.Table = NewTableService(s)
	s.Task = NewTaskService(s)
	s.Template = NewTemplateService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Column *ColumnService

	Query *QueryService

	Style *StyleService

	Table *TableService

	Task *TaskService

	Template *TemplateService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewColumnService(s *Service) *ColumnService {
	rs := &ColumnService{s: s}
	return rs
}

type ColumnService struct {
	s *Service
}

func NewQueryService(s *Service) *QueryService {
	rs := &QueryService{s: s}
	return rs
}

type QueryService struct {
	s *Service
}

func NewStyleService(s *Service) *StyleService {
	rs := &StyleService{s: s}
	return rs
}

type StyleService struct {
	s *Service
}

func NewTableService(s *Service) *TableService {
	rs := &TableService{s: s}
	return rs
}

type TableService struct {
	s *Service
}

func NewTaskService(s *Service) *TaskService {
	rs := &TaskService{s: s}
	return rs
}

type TaskService struct {
	s *Service
}

func NewTemplateService(s *Service) *TemplateService {
	rs := &TemplateService{s: s}
	return rs
}

type TemplateService struct {
	s *Service
}

// Bucket: Specifies the minimum and maximum values, the color, opacity,
// icon and weight of a bucket within a StyleSetting.
type Bucket struct {
	// Color: Color of line or the interior of a polygon in #RRGGBB format.
	Color string `json:"color,omitempty"`

	// Icon: Icon name used for a point.
	Icon string `json:"icon,omitempty"`

	// Max: Maximum value in the selected column for a row to be styled
	// according to the bucket color, opacity, icon, or weight.
	Max float64 `json:"max,omitempty"`

	// Min: Minimum value in the selected column for a row to be styled
	// according to the bucket color, opacity, icon, or weight.
	Min float64 `json:"min,omitempty"`

	// Opacity: Opacity of the color: 0.0 (transparent) to 1.0 (opaque).
	Opacity float64 `json:"opacity,omitempty"`

	// Weight: Width of a line (in pixels).
	Weight int64 `json:"weight,omitempty"`
}

// Column: Specifies the id, name and type of a column in a table.
type Column struct {
	// BaseColumn: Optional identifier of the base column. If present, this
	// column is derived from the specified base column.
	BaseColumn *ColumnBaseColumn `json:"baseColumn,omitempty"`

	// ColumnId: Identifier for the column.
	ColumnId int64 `json:"columnId,omitempty"`

	// Description: Optional column description.
	Description string `json:"description,omitempty"`

	// GraphPredicate: Optional column predicate. Used to map table to graph
	// data model (subject,predicate,object) See
	// http://www.w3.org/TR/2014/REC-rdf11-concepts-20140225/#data-model
	GraphPredicate string `json:"graph_predicate,omitempty"`

	// Kind: Type name: a template for an individual column.
	Kind string `json:"kind,omitempty"`

	// Name: Required name of the column.
	Name string `json:"name,omitempty"`

	// Type: Required type of the column.
	Type string `json:"type,omitempty"`
}

// ColumnBaseColumn: Optional identifier of the base column. If present,
// this column is derived from the specified base column.
type ColumnBaseColumn struct {
	// ColumnId: The id of the column in the base table from which this
	// column is derived.
	ColumnId int64 `json:"columnId,omitempty"`

	// TableIndex: Offset to the entry in the list of base tables in the
	// table definition.
	TableIndex int64 `json:"tableIndex,omitempty"`
}

// ColumnList: Represents a list of columns in a table.
type ColumnList struct {
	// Items: List of all requested columns.
	Items []*Column `json:"items,omitempty"`

	// Kind: Type name: a list of all columns.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of columns for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// Geometry: Represents a Geometry object.
type Geometry struct {
	// Geometries: The list of geometries in this geometry collection.
	Geometries []interface{} `json:"geometries,omitempty"`

	Geometry interface{} `json:"geometry,omitempty"`

	// Type: Type: A collection of geometries.
	Type string `json:"type,omitempty"`
}

// Import: Represents an import request.
type Import struct {
	// Kind: Type name: a template for an import request.
	Kind string `json:"kind,omitempty"`

	// NumRowsReceived: The number of rows received from the import request.
	NumRowsReceived int64 `json:"numRowsReceived,omitempty,string"`
}

// Line: Represents a line geometry.
type Line struct {
	// Coordinates: The coordinates that define the line.
	Coordinates [][]float64 `json:"coordinates,omitempty"`

	// Type: Type: A line geometry.
	Type string `json:"type,omitempty"`
}

// LineStyle: Represents a LineStyle within a StyleSetting
type LineStyle struct {
	// StrokeColor: Color of the line in #RRGGBB format.
	StrokeColor string `json:"strokeColor,omitempty"`

	// StrokeColorStyler: Column-value, gradient or buckets styler that is
	// used to determine the line color and opacity.
	StrokeColorStyler *StyleFunction `json:"strokeColorStyler,omitempty"`

	// StrokeOpacity: Opacity of the line : 0.0 (transparent) to 1.0
	// (opaque).
	StrokeOpacity float64 `json:"strokeOpacity,omitempty"`

	// StrokeWeight: Width of the line in pixels.
	StrokeWeight int64 `json:"strokeWeight,omitempty"`

	// StrokeWeightStyler: Column-value or bucket styler that is used to
	// determine the width of the line.
	StrokeWeightStyler *StyleFunction `json:"strokeWeightStyler,omitempty"`
}

// Point: Represents a point object.
type Point struct {
	// Coordinates: The coordinates that define the point.
	Coordinates []float64 `json:"coordinates,omitempty"`

	// Type: Point: A point geometry.
	Type string `json:"type,omitempty"`
}

// PointStyle: Represents a PointStyle within a StyleSetting
type PointStyle struct {
	// IconName: Name of the icon. Use values defined in
	// http://www.google.com/fusiontables/DataSource?dsrcid=308519
	IconName string `json:"iconName,omitempty"`

	// IconStyler: Column or a bucket value from which the icon name is to
	// be determined.
	IconStyler *StyleFunction `json:"iconStyler,omitempty"`
}

// Polygon: Represents a polygon object.
type Polygon struct {
	// Coordinates: The coordinates that define the polygon.
	Coordinates [][][]float64 `json:"coordinates,omitempty"`

	// Type: Type: A polygon geometry.
	Type string `json:"type,omitempty"`
}

// PolygonStyle: Represents a PolygonStyle within a StyleSetting
type PolygonStyle struct {
	// FillColor: Color of the interior of the polygon in #RRGGBB format.
	FillColor string `json:"fillColor,omitempty"`

	// FillColorStyler: Column-value, gradient, or bucket styler that is
	// used to determine the interior color and opacity of the polygon.
	FillColorStyler *StyleFunction `json:"fillColorStyler,omitempty"`

	// FillOpacity: Opacity of the interior of the polygon: 0.0
	// (transparent) to 1.0 (opaque).
	FillOpacity float64 `json:"fillOpacity,omitempty"`

	// StrokeColor: Color of the polygon border in #RRGGBB format.
	StrokeColor string `json:"strokeColor,omitempty"`

	// StrokeColorStyler: Column-value, gradient or buckets styler that is
	// used to determine the border color and opacity.
	StrokeColorStyler *StyleFunction `json:"strokeColorStyler,omitempty"`

	// StrokeOpacity: Opacity of the polygon border: 0.0 (transparent) to
	// 1.0 (opaque).
	StrokeOpacity float64 `json:"strokeOpacity,omitempty"`

	// StrokeWeight: Width of the polyon border in pixels.
	StrokeWeight int64 `json:"strokeWeight,omitempty"`

	// StrokeWeightStyler: Column-value or bucket styler that is used to
	// determine the width of the polygon border.
	StrokeWeightStyler *StyleFunction `json:"strokeWeightStyler,omitempty"`
}

// Sqlresponse: Represents a response to an sql statement.
type Sqlresponse struct {
	// Columns: Columns in the table.
	Columns []string `json:"columns,omitempty"`

	// Kind: Type name: a template for an individual table.
	Kind string `json:"kind,omitempty"`

	// Rows: The rows in the table. For each cell we print out whatever cell
	// value (e.g., numeric, string) exists. Thus it is important that each
	// cell contains only one value.
	Rows [][]interface{} `json:"rows,omitempty"`
}

// StyleFunction: Represents a StyleFunction within a StyleSetting
type StyleFunction struct {
	// Buckets: Bucket function that assigns a style based on the range a
	// column value falls into.
	Buckets []*Bucket `json:"buckets,omitempty"`

	// ColumnName: Name of the column whose value is used in the style.
	ColumnName string `json:"columnName,omitempty"`

	// Gradient: Gradient function that interpolates a range of colors based
	// on column value.
	Gradient *StyleFunctionGradient `json:"gradient,omitempty"`

	// Kind: Stylers can be one of three kinds: "fusiontables#fromColumn" if
	// the column value is to be used as is, i.e., the column values can
	// have colors in #RRGGBBAA format or integer line widths or icon names;
	// "fusiontables#gradient" if the styling of the row is to be based on
	// applying the gradient function on the column value; or
	// "fusiontables#buckets" if the styling is to based on the bucket into
	// which the the column value falls.
	Kind string `json:"kind,omitempty"`
}

// StyleFunctionGradient: Gradient function that interpolates a range of
// colors based on column value.
type StyleFunctionGradient struct {
	// Colors: Array with two or more colors.
	Colors []*StyleFunctionGradientColors `json:"colors,omitempty"`

	// Max: Higher-end of the interpolation range: rows with this value will
	// be assigned to colors[n-1].
	Max float64 `json:"max,omitempty"`

	// Min: Lower-end of the interpolation range: rows with this value will
	// be assigned to colors[0].
	Min float64 `json:"min,omitempty"`
}

type StyleFunctionGradientColors struct {
	// Color: Color in #RRGGBB format.
	Color string `json:"color,omitempty"`

	// Opacity: Opacity of the color: 0.0 (transparent) to 1.0 (opaque).
	Opacity float64 `json:"opacity,omitempty"`
}

// StyleSetting: Represents a complete StyleSettings object. The primary
// key is a combination of the tableId and a styleId.
type StyleSetting struct {
	// Kind: Type name: an individual style setting. A StyleSetting contains
	// the style defintions for points, lines, and polygons in a table.
	// Since a table can have any one or all of them, a style definition can
	// have point, line and polygon style definitions.
	Kind string `json:"kind,omitempty"`

	// MarkerOptions: Style definition for points in the table.
	MarkerOptions *PointStyle `json:"markerOptions,omitempty"`

	// Name: Optional name for the style setting.
	Name string `json:"name,omitempty"`

	// PolygonOptions: Style definition for polygons in the table.
	PolygonOptions *PolygonStyle `json:"polygonOptions,omitempty"`

	// PolylineOptions: Style definition for lines in the table.
	PolylineOptions *LineStyle `json:"polylineOptions,omitempty"`

	// StyleId: Identifier for the style setting (unique only within
	// tables).
	StyleId int64 `json:"styleId,omitempty"`

	// TableId: Identifier for the table.
	TableId string `json:"tableId,omitempty"`
}

// StyleSettingList: Represents a list of styles for a given table.
type StyleSettingList struct {
	// Items: All requested style settings.
	Items []*StyleSetting `json:"items,omitempty"`

	// Kind: Type name: in this case, a list of style settings.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of styles for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// Table: Represents a table. Specifies the name, whether it is
// exportable, description, attribution, and attribution link.
type Table struct {
	// Attribution: Optional attribution assigned to the table.
	Attribution string `json:"attribution,omitempty"`

	// AttributionLink: Optional link for attribution.
	AttributionLink string `json:"attributionLink,omitempty"`

	// BaseTableIds: Optional base table identifier if this table is a view
	// or merged table.
	BaseTableIds []string `json:"baseTableIds,omitempty"`

	// Columns: Columns in the table.
	Columns []*Column `json:"columns,omitempty"`

	// Description: Optional description assigned to the table.
	Description string `json:"description,omitempty"`

	// IsExportable: Variable for whether table is exportable.
	IsExportable bool `json:"isExportable,omitempty"`

	// Kind: Type name: a template for an individual table.
	Kind string `json:"kind,omitempty"`

	// Name: Name assigned to a table.
	Name string `json:"name,omitempty"`

	// Sql: Optional sql that encodes the table definition for derived
	// tables.
	Sql string `json:"sql,omitempty"`

	// TableId: Encrypted unique alphanumeric identifier for the table.
	TableId string `json:"tableId,omitempty"`
}

// TableList: Represents a list of tables.
type TableList struct {
	// Items: List of all requested tables.
	Items []*Table `json:"items,omitempty"`

	// Kind: Type name: a list of all tables.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// Task: Specifies the identifier, name, and type of a task in a table.
type Task struct {
	// Kind: Type of the resource. This is always "fusiontables#task".
	Kind string `json:"kind,omitempty"`

	// Progress: An indication of task progress.
	Progress string `json:"progress,omitempty"`

	// Started: false while the table is busy with some other task. true if
	// this background task is currently running.
	Started bool `json:"started,omitempty"`

	// TaskId: Identifier for the task.
	TaskId int64 `json:"taskId,omitempty,string"`

	// Type: Type of background task. One of  DELETE_ROWS Deletes one or
	// more rows from the table. ADD_ROWS "Adds one or more rows to a table.
	// Includes importing data into a new table and importing more rows into
	// an existing table. ADD_COLUMN Adds a new column to the table.
	// CHANGE_TYPE Changes the type of a column.
	Type string `json:"type,omitempty"`
}

// TaskList: Represents a list of tasks for a table.
type TaskList struct {
	// Items: List of all requested tasks.
	Items []*Task `json:"items,omitempty"`

	// Kind: Type of the resource. This is always "fusiontables#taskList".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of tasks for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// Template: Represents the contents of InfoWindow templates.
type Template struct {
	// AutomaticColumnNames: List of columns from which the template is to
	// be automatically constructed. Only one of body or automaticColumns
	// can be specified.
	AutomaticColumnNames []string `json:"automaticColumnNames,omitempty"`

	// Body: Body of the template. It contains HTML with {column_name} to
	// insert values from a particular column. The body is sanitized to
	// remove certain tags, e.g., script. Only one of body or
	// automaticColumns can be specified.
	Body string `json:"body,omitempty"`

	// Kind: Type name: a template for the info window contents. The
	// template can either include an HTML body or a list of columns from
	// which the template is computed automatically.
	Kind string `json:"kind,omitempty"`

	// Name: Optional name assigned to a template.
	Name string `json:"name,omitempty"`

	// TableId: Identifier for the table for which the template is defined.
	TableId string `json:"tableId,omitempty"`

	// TemplateId: Identifier for the template, unique within the context of
	// a particular table.
	TemplateId int64 `json:"templateId,omitempty"`
}

// TemplateList: Represents a list of templates for a given table.
type TemplateList struct {
	// Items: List of all requested templates.
	Items []*Template `json:"items,omitempty"`

	// Kind: Type name: a list of all templates.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of templates for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// method id "fusiontables.column.delete":

type ColumnDeleteCall struct {
	s        *Service
	tableId  string
	columnId string
	opt_     map[string]interface{}
}

// Delete: Deletes the column.
func (r *ColumnService) Delete(tableId string, columnId string) *ColumnDeleteCall {
	c := &ColumnDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.columnId = columnId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnDeleteCall) Fields(s ...googleapi.Field) *ColumnDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ColumnDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/columns/{columnId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ColumnDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes the column.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.column.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table from which the column is being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.column.get":

type ColumnGetCall struct {
	s        *Service
	tableId  string
	columnId string
	opt_     map[string]interface{}
}

// Get: Retrieves a specific column by its id.
func (r *ColumnService) Get(tableId string, columnId string) *ColumnGetCall {
	c := &ColumnGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.columnId = columnId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnGetCall) Fields(s ...googleapi.Field) *ColumnGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ColumnGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/columns/{columnId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ColumnGetCall) Do() (*Column, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Column
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a specific column by its id.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.column.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column that is being requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table to which the column belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.column.insert":

type ColumnInsertCall struct {
	s       *Service
	tableId string
	column  *Column
	opt_    map[string]interface{}
}

// Insert: Adds a new column to the table.
func (r *ColumnService) Insert(tableId string, column *Column) *ColumnInsertCall {
	c := &ColumnInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.column = column
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnInsertCall) Fields(s ...googleapi.Field) *ColumnInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ColumnInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.column)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/columns")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ColumnInsertCall) Do() (*Column, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Column
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a new column to the table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.column.insert",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table for which a new column is being added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns",
	//   "request": {
	//     "$ref": "Column"
	//   },
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.column.list":

type ColumnListCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// List: Retrieves a list of columns.
func (r *ColumnService) List(tableId string) *ColumnListCall {
	c := &ColumnListCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of columns to return.  Default is 5.
func (c *ColumnListCall) MaxResults(maxResults int64) *ColumnListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *ColumnListCall) PageToken(pageToken string) *ColumnListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnListCall) Fields(s ...googleapi.Field) *ColumnListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ColumnListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/columns")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ColumnListCall) Do() (*ColumnList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ColumnList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of columns.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.column.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of columns to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table whose columns are being listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns",
	//   "response": {
	//     "$ref": "ColumnList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.column.patch":

type ColumnPatchCall struct {
	s        *Service
	tableId  string
	columnId string
	column   *Column
	opt_     map[string]interface{}
}

// Patch: Updates the name or type of an existing column. This method
// supports patch semantics.
func (r *ColumnService) Patch(tableId string, columnId string, column *Column) *ColumnPatchCall {
	c := &ColumnPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.columnId = columnId
	c.column = column
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnPatchCall) Fields(s ...googleapi.Field) *ColumnPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ColumnPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.column)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/columns/{columnId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ColumnPatchCall) Do() (*Column, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Column
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the name or type of an existing column. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.column.patch",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table for which the column is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "request": {
	//     "$ref": "Column"
	//   },
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.column.update":

type ColumnUpdateCall struct {
	s        *Service
	tableId  string
	columnId string
	column   *Column
	opt_     map[string]interface{}
}

// Update: Updates the name or type of an existing column.
func (r *ColumnService) Update(tableId string, columnId string, column *Column) *ColumnUpdateCall {
	c := &ColumnUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.columnId = columnId
	c.column = column
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnUpdateCall) Fields(s ...googleapi.Field) *ColumnUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ColumnUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.column)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/columns/{columnId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *ColumnUpdateCall) Do() (*Column, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Column
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the name or type of an existing column.",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.column.update",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table for which the column is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "request": {
	//     "$ref": "Column"
	//   },
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.query.sql":

type QuerySqlCall struct {
	s    *Service
	sql  string
	opt_ map[string]interface{}
}

// Sql: Executes an SQL SELECT/INSERT/UPDATE/DELETE/SHOW/DESCRIBE/CREATE
// statement.
func (r *QueryService) Sql(sql string) *QuerySqlCall {
	c := &QuerySqlCall{s: r.s, opt_: make(map[string]interface{})}
	c.sql = sql
	return c
}

// Hdrs sets the optional parameter "hdrs": Should column names be
// included (in the first row)?. Default is true.
func (c *QuerySqlCall) Hdrs(hdrs bool) *QuerySqlCall {
	c.opt_["hdrs"] = hdrs
	return c
}

// Typed sets the optional parameter "typed": Should typed values be
// returned in the (JSON) response -- numbers for numeric values and
// parsed geometries for KML values? Default is true.
func (c *QuerySqlCall) Typed(typed bool) *QuerySqlCall {
	c.opt_["typed"] = typed
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySqlCall) Fields(s ...googleapi.Field) *QuerySqlCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *QuerySqlCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("sql", fmt.Sprintf("%v", c.sql))
	if v, ok := c.opt_["hdrs"]; ok {
		params.Set("hdrs", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["typed"]; ok {
		params.Set("typed", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "query")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

// Download fetches the API endpoint's "media" value, instead of the normal
// API response value. If the returned error is nil, the Response is guaranteed to
// have a 2xx status code. Callers must close the Response.Body as usual.
func (c *QuerySqlCall) Download() (*http.Response, error) {
	res, err := c.doRequest("media")
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckMediaResponse(res); err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, nil
}

func (c *QuerySqlCall) Do() (*Sqlresponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Sqlresponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Executes an SQL SELECT/INSERT/UPDATE/DELETE/SHOW/DESCRIBE/CREATE statement.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.query.sql",
	//   "parameterOrder": [
	//     "sql"
	//   ],
	//   "parameters": {
	//     "hdrs": {
	//       "description": "Should column names be included (in the first row)?. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sql": {
	//       "description": "An SQL SELECT/SHOW/DESCRIBE/INSERT/UPDATE/DELETE/CREATE statement.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typed": {
	//       "description": "Should typed values be returned in the (JSON) response -- numbers for numeric values and parsed geometries for KML values? Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "query",
	//   "response": {
	//     "$ref": "Sqlresponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "fusiontables.query.sqlGet":

type QuerySqlGetCall struct {
	s    *Service
	sql  string
	opt_ map[string]interface{}
}

// SqlGet: Executes an SQL SELECT/SHOW/DESCRIBE statement.
func (r *QueryService) SqlGet(sql string) *QuerySqlGetCall {
	c := &QuerySqlGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.sql = sql
	return c
}

// Hdrs sets the optional parameter "hdrs": Should column names be
// included (in the first row)?. Default is true.
func (c *QuerySqlGetCall) Hdrs(hdrs bool) *QuerySqlGetCall {
	c.opt_["hdrs"] = hdrs
	return c
}

// Typed sets the optional parameter "typed": Should typed values be
// returned in the (JSON) response -- numbers for numeric values and
// parsed geometries for KML values? Default is true.
func (c *QuerySqlGetCall) Typed(typed bool) *QuerySqlGetCall {
	c.opt_["typed"] = typed
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySqlGetCall) Fields(s ...googleapi.Field) *QuerySqlGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *QuerySqlGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("sql", fmt.Sprintf("%v", c.sql))
	if v, ok := c.opt_["hdrs"]; ok {
		params.Set("hdrs", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["typed"]; ok {
		params.Set("typed", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "query")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

// Download fetches the API endpoint's "media" value, instead of the normal
// API response value. If the returned error is nil, the Response is guaranteed to
// have a 2xx status code. Callers must close the Response.Body as usual.
func (c *QuerySqlGetCall) Download() (*http.Response, error) {
	res, err := c.doRequest("media")
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckMediaResponse(res); err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, nil
}

func (c *QuerySqlGetCall) Do() (*Sqlresponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Sqlresponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Executes an SQL SELECT/SHOW/DESCRIBE statement.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.query.sqlGet",
	//   "parameterOrder": [
	//     "sql"
	//   ],
	//   "parameters": {
	//     "hdrs": {
	//       "description": "Should column names be included (in the first row)?. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sql": {
	//       "description": "An SQL SELECT/SHOW/DESCRIBE statement.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typed": {
	//       "description": "Should typed values be returned in the (JSON) response -- numbers for numeric values and parsed geometries for KML values? Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "query",
	//   "response": {
	//     "$ref": "Sqlresponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "fusiontables.style.delete":

type StyleDeleteCall struct {
	s       *Service
	tableId string
	styleId int64
	opt_    map[string]interface{}
}

// Delete: Deletes a style.
func (r *StyleService) Delete(tableId string, styleId int64) *StyleDeleteCall {
	c := &StyleDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.styleId = styleId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleDeleteCall) Fields(s ...googleapi.Field) *StyleDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StyleDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/styles/{styleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *StyleDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a style.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.style.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (within a table) for the style being deleted",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table from which the style is being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.style.get":

type StyleGetCall struct {
	s       *Service
	tableId string
	styleId int64
	opt_    map[string]interface{}
}

// Get: Gets a specific style.
func (r *StyleService) Get(tableId string, styleId int64) *StyleGetCall {
	c := &StyleGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.styleId = styleId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleGetCall) Fields(s ...googleapi.Field) *StyleGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StyleGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/styles/{styleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *StyleGetCall) Do() (*StyleSetting, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *StyleSetting
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a specific style.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.style.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (integer) for a specific style in a table",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table to which the requested style belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.style.insert":

type StyleInsertCall struct {
	s            *Service
	tableId      string
	stylesetting *StyleSetting
	opt_         map[string]interface{}
}

// Insert: Adds a new style for the table.
func (r *StyleService) Insert(tableId string, stylesetting *StyleSetting) *StyleInsertCall {
	c := &StyleInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.stylesetting = stylesetting
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleInsertCall) Fields(s ...googleapi.Field) *StyleInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StyleInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.stylesetting)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/styles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *StyleInsertCall) Do() (*StyleSetting, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *StyleSetting
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds a new style for the table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.style.insert",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table for which a new style is being added",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles",
	//   "request": {
	//     "$ref": "StyleSetting"
	//   },
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.style.list":

type StyleListCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// List: Retrieves a list of styles.
func (r *StyleService) List(tableId string) *StyleListCall {
	c := &StyleListCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of styles to return.  Default is 5.
func (c *StyleListCall) MaxResults(maxResults int64) *StyleListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *StyleListCall) PageToken(pageToken string) *StyleListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleListCall) Fields(s ...googleapi.Field) *StyleListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StyleListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/styles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *StyleListCall) Do() (*StyleSettingList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *StyleSettingList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of styles.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.style.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of styles to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table whose styles are being listed",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles",
	//   "response": {
	//     "$ref": "StyleSettingList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.style.patch":

type StylePatchCall struct {
	s            *Service
	tableId      string
	styleId      int64
	stylesetting *StyleSetting
	opt_         map[string]interface{}
}

// Patch: Updates an existing style. This method supports patch
// semantics.
func (r *StyleService) Patch(tableId string, styleId int64, stylesetting *StyleSetting) *StylePatchCall {
	c := &StylePatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.styleId = styleId
	c.stylesetting = stylesetting
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StylePatchCall) Fields(s ...googleapi.Field) *StylePatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StylePatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.stylesetting)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/styles/{styleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *StylePatchCall) Do() (*StyleSetting, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *StyleSetting
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing style. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.style.patch",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (within a table) for the style being updated.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose style is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "request": {
	//     "$ref": "StyleSetting"
	//   },
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.style.update":

type StyleUpdateCall struct {
	s            *Service
	tableId      string
	styleId      int64
	stylesetting *StyleSetting
	opt_         map[string]interface{}
}

// Update: Updates an existing style.
func (r *StyleService) Update(tableId string, styleId int64, stylesetting *StyleSetting) *StyleUpdateCall {
	c := &StyleUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.styleId = styleId
	c.stylesetting = stylesetting
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleUpdateCall) Fields(s ...googleapi.Field) *StyleUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *StyleUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.stylesetting)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/styles/{styleId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *StyleUpdateCall) Do() (*StyleSetting, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *StyleSetting
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing style.",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.style.update",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (within a table) for the style being updated.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose style is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "request": {
	//     "$ref": "StyleSetting"
	//   },
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.copy":

type TableCopyCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// Copy: Copies a table.
func (r *TableService) Copy(tableId string) *TableCopyCall {
	c := &TableCopyCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// CopyPresentation sets the optional parameter "copyPresentation":
// Whether to also copy tabs, styles, and templates. Default is false.
func (c *TableCopyCall) CopyPresentation(copyPresentation bool) *TableCopyCall {
	c.opt_["copyPresentation"] = copyPresentation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableCopyCall) Fields(s ...googleapi.Field) *TableCopyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableCopyCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["copyPresentation"]; ok {
		params.Set("copyPresentation", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/copy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableCopyCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Copies a table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.copy",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "copyPresentation": {
	//       "description": "Whether to also copy tabs, styles, and templates. Default is false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "tableId": {
	//       "description": "ID of the table that is being copied.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/copy",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.table.delete":

type TableDeleteCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// Delete: Deletes a table.
func (r *TableService) Delete(tableId string) *TableDeleteCall {
	c := &TableDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableDeleteCall) Fields(s ...googleapi.Field) *TableDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a table.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.table.delete",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "ID of the table that is being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.get":

type TableGetCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// Get: Retrieves a specific table by its id.
func (r *TableService) Get(tableId string) *TableGetCall {
	c := &TableGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableGetCall) Fields(s ...googleapi.Field) *TableGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableGetCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a specific table by its id.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.table.get",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Identifier(ID) for the table being requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.table.importRows":

type TableImportRowsCall struct {
	s          *Service
	tableId    string
	opt_       map[string]interface{}
	media_     io.Reader
	resumable_ googleapi.SizeReaderAt
	mediaType_ string
	ctx_       context.Context
	protocol_  string
}

// ImportRows: Import more rows into a table.
func (r *TableService) ImportRows(tableId string) *TableImportRowsCall {
	c := &TableImportRowsCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// Delimiter sets the optional parameter "delimiter": The delimiter used
// to separate cell values. This can only consist of a single character.
// Default is ','.
func (c *TableImportRowsCall) Delimiter(delimiter string) *TableImportRowsCall {
	c.opt_["delimiter"] = delimiter
	return c
}

// Encoding sets the optional parameter "encoding": The encoding of the
// content. Default is UTF-8. Use 'auto-detect' if you are unsure of the
// encoding.
func (c *TableImportRowsCall) Encoding(encoding string) *TableImportRowsCall {
	c.opt_["encoding"] = encoding
	return c
}

// EndLine sets the optional parameter "endLine": The index of the last
// line from which to start importing, exclusive. Thus, the number of
// imported lines is endLine - startLine. If this parameter is not
// provided, the file will be imported until the last line of the file.
// If endLine is negative, then the imported content will exclude the
// last endLine lines. That is, if endline is negative, no line will be
// imported whose index is greater than N + endLine where N is the
// number of lines in the file, and the number of imported lines will be
// N + endLine - startLine.
func (c *TableImportRowsCall) EndLine(endLine int64) *TableImportRowsCall {
	c.opt_["endLine"] = endLine
	return c
}

// IsStrict sets the optional parameter "isStrict": Whether the CSV must
// have the same number of values for each row. If false, rows with
// fewer values will be padded with empty values. Default is true.
func (c *TableImportRowsCall) IsStrict(isStrict bool) *TableImportRowsCall {
	c.opt_["isStrict"] = isStrict
	return c
}

// StartLine sets the optional parameter "startLine": The index of the
// first line from which to start importing, inclusive. Default is 0.
func (c *TableImportRowsCall) StartLine(startLine int64) *TableImportRowsCall {
	c.opt_["startLine"] = startLine
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
func (c *TableImportRowsCall) Media(r io.Reader) *TableImportRowsCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TableImportRowsCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TableImportRowsCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TableImportRowsCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TableImportRowsCall {
	c.opt_["progressUpdater"] = pu
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableImportRowsCall) Fields(s ...googleapi.Field) *TableImportRowsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableImportRowsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["delimiter"]; ok {
		params.Set("delimiter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["encoding"]; ok {
		params.Set("encoding", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["endLine"]; ok {
		params.Set("endLine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["isStrict"]; ok {
		params.Set("isStrict", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startLine"]; ok {
		params.Set("startLine", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/import")
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		params.Set("uploadType", c.protocol_)
	}
	urls += "?" + params.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableImportRowsCall) Do() (*Import, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var progressUpdater_ googleapi.ProgressUpdater
	if v, ok := c.opt_["progressUpdater"]; ok {
		if pu, ok := v.(googleapi.ProgressUpdater); ok {
			progressUpdater_ = pu
		}
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	var ret *Import
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Import more rows into a table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.importRows",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "maxSize": "250MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/fusiontables/v1/tables/{tableId}/import"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/fusiontables/v1/tables/{tableId}/import"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "delimiter": {
	//       "description": "The delimiter used to separate cell values. This can only consist of a single character. Default is ','.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "encoding": {
	//       "description": "The encoding of the content. Default is UTF-8. Use 'auto-detect' if you are unsure of the encoding.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endLine": {
	//       "description": "The index of the last line from which to start importing, exclusive. Thus, the number of imported lines is endLine - startLine. If this parameter is not provided, the file will be imported until the last line of the file. If endLine is negative, then the imported content will exclude the last endLine lines. That is, if endline is negative, no line will be imported whose index is greater than N + endLine where N is the number of lines in the file, and the number of imported lines will be N + endLine - startLine.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "isStrict": {
	//       "description": "Whether the CSV must have the same number of values for each row. If false, rows with fewer values will be padded with empty values. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "startLine": {
	//       "description": "The index of the first line from which to start importing, inclusive. Default is 0.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "The table into which new rows are being imported.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/import",
	//   "response": {
	//     "$ref": "Import"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "fusiontables.table.importTable":

type TableImportTableCall struct {
	s          *Service
	name       string
	opt_       map[string]interface{}
	media_     io.Reader
	resumable_ googleapi.SizeReaderAt
	mediaType_ string
	ctx_       context.Context
	protocol_  string
}

// ImportTable: Import a new table.
func (r *TableService) ImportTable(name string) *TableImportTableCall {
	c := &TableImportTableCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Delimiter sets the optional parameter "delimiter": The delimiter used
// to separate cell values. This can only consist of a single character.
// Default is ','.
func (c *TableImportTableCall) Delimiter(delimiter string) *TableImportTableCall {
	c.opt_["delimiter"] = delimiter
	return c
}

// Encoding sets the optional parameter "encoding": The encoding of the
// content. Default is UTF-8. Use 'auto-detect' if you are unsure of the
// encoding.
func (c *TableImportTableCall) Encoding(encoding string) *TableImportTableCall {
	c.opt_["encoding"] = encoding
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
func (c *TableImportTableCall) Media(r io.Reader) *TableImportTableCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TableImportTableCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TableImportTableCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TableImportTableCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TableImportTableCall {
	c.opt_["progressUpdater"] = pu
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableImportTableCall) Fields(s ...googleapi.Field) *TableImportTableCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableImportTableCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("name", fmt.Sprintf("%v", c.name))
	if v, ok := c.opt_["delimiter"]; ok {
		params.Set("delimiter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["encoding"]; ok {
		params.Set("encoding", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/import")
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		params.Set("uploadType", c.protocol_)
	}
	urls += "?" + params.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableImportTableCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var progressUpdater_ googleapi.ProgressUpdater
	if v, ok := c.opt_["progressUpdater"]; ok {
		if pu, ok := v.(googleapi.ProgressUpdater); ok {
			progressUpdater_ = pu
		}
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Import a new table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.importTable",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "maxSize": "250MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/fusiontables/v1/tables/import"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/fusiontables/v1/tables/import"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "delimiter": {
	//       "description": "The delimiter used to separate cell values. This can only consist of a single character. Default is ','.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "encoding": {
	//       "description": "The encoding of the content. Default is UTF-8. Use 'auto-detect' if you are unsure of the encoding.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name to be assigned to the new table.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/import",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "fusiontables.table.insert":

type TableInsertCall struct {
	s     *Service
	table *Table
	opt_  map[string]interface{}
}

// Insert: Creates a new table.
func (r *TableService) Insert(table *Table) *TableInsertCall {
	c := &TableInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.table = table
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableInsertCall) Fields(s ...googleapi.Field) *TableInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableInsertCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.insert",
	//   "path": "tables",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.list":

type TableListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Retrieves a list of tables a user owns.
func (r *TableService) List() *TableListCall {
	c := &TableListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of styles to return.  Default is 5.
func (c *TableListCall) MaxResults(maxResults int64) *TableListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *TableListCall) PageToken(pageToken string) *TableListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableListCall) Fields(s ...googleapi.Field) *TableListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableListCall) Do() (*TableList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TableList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of tables a user owns.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.table.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of styles to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables",
	//   "response": {
	//     "$ref": "TableList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.table.patch":

type TablePatchCall struct {
	s       *Service
	tableId string
	table   *Table
	opt_    map[string]interface{}
}

// Patch: Updates an existing table. Unless explicitly requested, only
// the name, description, and attribution will be updated. This method
// supports patch semantics.
func (r *TableService) Patch(tableId string, table *Table) *TablePatchCall {
	c := &TablePatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.table = table
	return c
}

// ReplaceViewDefinition sets the optional parameter
// "replaceViewDefinition": Should the view definition also be updated?
// The specified view definition replaces the existing one. Only a view
// can be updated with a new definition.
func (c *TablePatchCall) ReplaceViewDefinition(replaceViewDefinition bool) *TablePatchCall {
	c.opt_["replaceViewDefinition"] = replaceViewDefinition
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablePatchCall) Fields(s ...googleapi.Field) *TablePatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TablePatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["replaceViewDefinition"]; ok {
		params.Set("replaceViewDefinition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TablePatchCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing table. Unless explicitly requested, only the name, description, and attribution will be updated. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.table.patch",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "replaceViewDefinition": {
	//       "description": "Should the view definition also be updated? The specified view definition replaces the existing one. Only a view can be updated with a new definition.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "tableId": {
	//       "description": "ID of the table that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.update":

type TableUpdateCall struct {
	s       *Service
	tableId string
	table   *Table
	opt_    map[string]interface{}
}

// Update: Updates an existing table. Unless explicitly requested, only
// the name, description, and attribution will be updated.
func (r *TableService) Update(tableId string, table *Table) *TableUpdateCall {
	c := &TableUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.table = table
	return c
}

// ReplaceViewDefinition sets the optional parameter
// "replaceViewDefinition": Should the view definition also be updated?
// The specified view definition replaces the existing one. Only a view
// can be updated with a new definition.
func (c *TableUpdateCall) ReplaceViewDefinition(replaceViewDefinition bool) *TableUpdateCall {
	c.opt_["replaceViewDefinition"] = replaceViewDefinition
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableUpdateCall) Fields(s ...googleapi.Field) *TableUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TableUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["replaceViewDefinition"]; ok {
		params.Set("replaceViewDefinition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TableUpdateCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing table. Unless explicitly requested, only the name, description, and attribution will be updated.",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.table.update",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "replaceViewDefinition": {
	//       "description": "Should the view definition also be updated? The specified view definition replaces the existing one. Only a view can be updated with a new definition.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "tableId": {
	//       "description": "ID of the table that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.task.delete":

type TaskDeleteCall struct {
	s       *Service
	tableId string
	taskId  string
	opt_    map[string]interface{}
}

// Delete: Deletes the task, unless already started.
func (r *TaskService) Delete(tableId string, taskId string) *TaskDeleteCall {
	c := &TaskDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.taskId = taskId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TaskDeleteCall) Fields(s ...googleapi.Field) *TaskDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TaskDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/tasks/{taskId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"taskId":  c.taskId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TaskDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes the task, unless already started.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.task.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "taskId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table from which the task is being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/tasks/{taskId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.task.get":

type TaskGetCall struct {
	s       *Service
	tableId string
	taskId  string
	opt_    map[string]interface{}
}

// Get: Retrieves a specific task by its id.
func (r *TaskService) Get(tableId string, taskId string) *TaskGetCall {
	c := &TaskGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.taskId = taskId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TaskGetCall) Fields(s ...googleapi.Field) *TaskGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TaskGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/tasks/{taskId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"taskId":  c.taskId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TaskGetCall) Do() (*Task, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Task
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a specific task by its id.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.task.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "taskId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the task belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/tasks/{taskId}",
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.task.list":

type TaskListCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// List: Retrieves a list of tasks.
func (r *TaskService) List(tableId string) *TaskListCall {
	c := &TaskListCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of columns to return.  Default is 5.
func (c *TaskListCall) MaxResults(maxResults int64) *TaskListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *TaskListCall) PageToken(pageToken string) *TaskListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartIndex sets the optional parameter "startIndex":
func (c *TaskListCall) StartIndex(startIndex int64) *TaskListCall {
	c.opt_["startIndex"] = startIndex
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TaskListCall) Fields(s ...googleapi.Field) *TaskListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TaskListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startIndex"]; ok {
		params.Set("startIndex", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/tasks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TaskListCall) Do() (*TaskList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TaskList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of tasks.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.task.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of columns to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose tasks are being listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/tasks",
	//   "response": {
	//     "$ref": "TaskList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.template.delete":

type TemplateDeleteCall struct {
	s          *Service
	tableId    string
	templateId int64
	opt_       map[string]interface{}
}

// Delete: Deletes a template
func (r *TemplateService) Delete(tableId string, templateId int64) *TemplateDeleteCall {
	c := &TemplateDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.templateId = templateId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateDeleteCall) Fields(s ...googleapi.Field) *TemplateDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TemplateDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/templates/{templateId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TemplateDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a template",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.template.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table from which the template is being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template which is being deleted",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.template.get":

type TemplateGetCall struct {
	s          *Service
	tableId    string
	templateId int64
	opt_       map[string]interface{}
}

// Get: Retrieves a specific template by its id
func (r *TemplateService) Get(tableId string, templateId int64) *TemplateGetCall {
	c := &TemplateGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.templateId = templateId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateGetCall) Fields(s ...googleapi.Field) *TemplateGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TemplateGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/templates/{templateId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TemplateGetCall) Do() (*Template, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Template
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a specific template by its id",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.template.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the template belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template that is being requested",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.template.insert":

type TemplateInsertCall struct {
	s        *Service
	tableId  string
	template *Template
	opt_     map[string]interface{}
}

// Insert: Creates a new template for the table.
func (r *TemplateService) Insert(tableId string, template *Template) *TemplateInsertCall {
	c := &TemplateInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.template = template
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateInsertCall) Fields(s ...googleapi.Field) *TemplateInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TemplateInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.template)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/templates")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TemplateInsertCall) Do() (*Template, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Template
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new template for the table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.template.insert",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table for which a new template is being created",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates",
	//   "request": {
	//     "$ref": "Template"
	//   },
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.template.list":

type TemplateListCall struct {
	s       *Service
	tableId string
	opt_    map[string]interface{}
}

// List: Retrieves a list of templates.
func (r *TemplateService) List(tableId string) *TemplateListCall {
	c := &TemplateListCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of templates to return.  Default is 5.
func (c *TemplateListCall) MaxResults(maxResults int64) *TemplateListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which results page to return.
func (c *TemplateListCall) PageToken(pageToken string) *TemplateListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateListCall) Fields(s ...googleapi.Field) *TemplateListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TemplateListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/templates")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TemplateListCall) Do() (*TemplateList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TemplateList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a list of templates.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.template.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of templates to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which results page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Identifier for the table whose templates are being requested",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates",
	//   "response": {
	//     "$ref": "TemplateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.template.patch":

type TemplatePatchCall struct {
	s          *Service
	tableId    string
	templateId int64
	template   *Template
	opt_       map[string]interface{}
}

// Patch: Updates an existing template. This method supports patch
// semantics.
func (r *TemplateService) Patch(tableId string, templateId int64, template *Template) *TemplatePatchCall {
	c := &TemplatePatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.templateId = templateId
	c.template = template
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplatePatchCall) Fields(s ...googleapi.Field) *TemplatePatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TemplatePatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.template)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/templates/{templateId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TemplatePatchCall) Do() (*Template, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Template
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing template. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.template.patch",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the updated template belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template that is being updated",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "request": {
	//     "$ref": "Template"
	//   },
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.template.update":

type TemplateUpdateCall struct {
	s          *Service
	tableId    string
	templateId int64
	template   *Template
	opt_       map[string]interface{}
}

// Update: Updates an existing template
func (r *TemplateService) Update(tableId string, templateId int64, template *Template) *TemplateUpdateCall {
	c := &TemplateUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.tableId = tableId
	c.templateId = templateId
	c.template = template
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateUpdateCall) Fields(s ...googleapi.Field) *TemplateUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TemplateUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.template)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/templates/{templateId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	return c.s.client.Do(req)
}

func (c *TemplateUpdateCall) Do() (*Template, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Template
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing template",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.template.update",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the updated template belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template that is being updated",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "request": {
	//     "$ref": "Template"
	//   },
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}
