package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// TableShardResultsItem is a nested struct in drds response
type TableShardResultsItem struct {
	Table          string `json:"Table" xml:"Table"`
	ShardType      string `json:"ShardType" xml:"ShardType"`
	DbShardKey     string `json:"DbShardKey" xml:"DbShardKey"`
	TbShardKey     string `json:"TbShardKey" xml:"TbShardKey"`
	Tbpartitions   int    `json:"Tbpartitions" xml:"Tbpartitions"`
	Dbpartitions   int    `json:"Dbpartitions" xml:"Dbpartitions"`
	CreateTableSql string `json:"CreateTableSql" xml:"CreateTableSql"`
	RowCount       int    `json:"RowCount" xml:"RowCount"`
	SqlCount       int    `json:"SqlCount" xml:"SqlCount"`
	ShardKeyType   string `json:"ShardKeyType" xml:"ShardKeyType"`
}
