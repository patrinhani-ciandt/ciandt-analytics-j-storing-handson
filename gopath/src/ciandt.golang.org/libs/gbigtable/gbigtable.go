package gbigtable

import (
	"fmt"
	"time"
	"io/ioutil"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/bigtable"	
)

type ProcessDataTableRow func(DtRow)

type ClientConnectionData struct { 
	Project string
	Zone string
	Cluster string
	KeyJsonFilePath string
}

type DtRow struct {
	Key string
	Families map[string]map[string]interface{}
}

func getClientOptionFromJsonKeyFile(ctx context.Context, keyFilePath string, scope string) (cloud.ClientOption) {

	fmt.Println("Reading [" , keyFilePath, "] ...")
	jsonKey, err := ioutil.ReadFile(keyFilePath)
	
	if err != nil {
		fmt.Println("Error on [ioutil.ReadFile]: %v", err)
	}
	
	fmt.Println("Creating config ...")
	config, err := google.JWTConfigFromJSON(jsonKey, scope)

	if err != nil {
		fmt.Println("Error on [google.JWTConfigFromJSON]]: %v", err)
	}
	
	clientOption := cloud.WithTokenSource(config.TokenSource(ctx))
	
	return clientOption
}

func GetContext(timeout time.Duration) (context.Context) {
	
	ctx, _ := context.WithTimeout(context.Background(), timeout)

    return ctx
}

func OpenAdminClient(ctx context.Context, connectionData ClientConnectionData) (*bigtable.AdminClient) {

	clientOpt := getClientOptionFromJsonKeyFile(ctx, connectionData.KeyJsonFilePath, bigtable.AdminScope)
		
	client, err := bigtable.NewAdminClient(ctx, connectionData.Project, connectionData.Zone, connectionData.Cluster, clientOpt)	
	
	if err != nil {
		fmt.Println("Error on [NewAdminClient]: %v", err)
	}

	return client
}

func OpenClient(ctx context.Context, connectionData ClientConnectionData) (*bigtable.Client) {
	
	clientOpt := getClientOptionFromJsonKeyFile(ctx, connectionData.KeyJsonFilePath, bigtable.Scope)
		
	client, err := bigtable.NewClient(ctx, connectionData.Project, connectionData.Zone, connectionData.Cluster, clientOpt)	
	
	if err != nil {
		fmt.Println("Error on [NewClient]: %v", err)
	}

	return client
}

func DeleteTable(connectionData ClientConnectionData, ctx context.Context, tableName string) {

	adminClient := OpenAdminClient(ctx, connectionData)

	fmt.Println("Deleting table ...")
	err := adminClient.DeleteTable(ctx, tableName)
	if err != nil {
		fmt.Println("Error on [DeleteTable]: %v", err)
	}
	
	defer adminClient.Close()
}

func CreateTable(connectionData ClientConnectionData, ctx context.Context, tableName string, families []string) {
	
	adminClient := OpenAdminClient(ctx, connectionData)

	fmt.Println("Creating table ...")
	err := adminClient.CreateTable(ctx, tableName)
	if err != nil {
		fmt.Println("Error on [CreateTable]: %v", err)
	}

	fmt.Println("Creating families ...")	
	for i := 0; i < len(families); i++ {
		
		fmt.Println("Creating family: %v", families[i])

		err = adminClient.CreateColumnFamily(ctx, tableName, families[i])
		if err != nil {
			fmt.Println("Error on [CreateColumnFamily]: %v", err)
		}
	}
	
	defer adminClient.Close()
}

func OpenTable(table string, client *bigtable.Client) (*bigtable.Table) {

	fmt.Println("Opening Table ...")
	tbl := client.Open(table)
	
	return tbl
}

func WriteRow(ctx context.Context, table *bigtable.Table, columnFamilySep string, rowKey string, columns []string, rowCells []string, startCellIndex int) {

	mut := bigtable.NewMutation()

	for i := startCellIndex; i < len(rowCells); i++ {

		var colSet = strings.Split(columns[i], columnFamilySep)
		fam := colSet[0]
		col := colSet[1]
		  
		mut.Set(fam, col, 0, []byte(rowCells[i]))
	}
				
	fmt.Println("Applying row: ", rowKey)
	if err := table.Apply(ctx, rowKey, mut); err != nil {
		
		fmt.Println("Error on Mutating row %v: %v", rowKey, err)
	}
}

func extractDtRowFromBigTableRow(r bigtable.Row) DtRow {

	row := DtRow {
		Families: make(map[string]map[string]interface{}),
	}

	for _, ris := range r {

		for _, ri := range ris {
			
			var colSet = strings.Split(ri.Column, ":")
			fam := colSet[0]
			col := colSet[1]
		
			cellValue := fmt.Sprintf("%s", ri.Value)
			
			row.Key = ri.Row;
			
			mFamilies, okMapFamilies := row.Families[fam]
			if !okMapFamilies {
				mFamilies = make(map[string]interface{})
				row.Families[fam] = mFamilies
			}
			
			row.Families[fam][col] = cellValue;
		}
		
	}

	return row;
}

func ReadRow(ctx context.Context, table *bigtable.Table, rowKey string) DtRow {

	var row DtRow; 

	r, err := table.ReadRow(ctx, rowKey)

	if err != nil {
		
		fmt.Println("Error on [ReadRow]: %v", err)
	} else {
		
		row = extractDtRowFromBigTableRow(r)
	}
	
	return row;
}

func ReadRows(ctx context.Context, table *bigtable.Table, rowRange bigtable.RowRange, processDataTableRow ProcessDataTableRow, opts ...bigtable.ReadOption) {
	
	fmt.Println("Reading Rows ...")
	table.ReadRows(ctx, rowRange, func(r bigtable.Row) bool {
    	
		row := extractDtRowFromBigTableRow(r)
		
		processDataTableRow(row)

		return true
	}, opts...)
}

func ReadAllRows(ctx context.Context, table *bigtable.Table, processDataTableRow ProcessDataTableRow) {
	
	ReadRows(ctx, table, bigtable.InfiniteRange(""), processDataTableRow)
}