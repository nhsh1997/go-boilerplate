package apps



type ILogger interface {

}



/*type IRepository interface {
paginate(criteria: ICriteria): Promise<IPaginateResult<Entity>>;
find(criteria: ICriteria): Promise<Entity[]>;
findById(id: number, criteria?: ICriteria): Promise<Entity>;
create(entity: Entity, options?: { include?: any }): Promise<Entity>;
updateById(id: number, entity: any): Promise<Entity>;
deleteById(id: number): Promise<undefined>;
bulkCreate(entities: Entity[]): Promise<Entity[]>;
bulkUpdate(criteria: ICriteria, entities: any): Promise<Entity[]>;
bulkDelete(criteria: ICriteria): Promise<undefined>;
count(criteria: ICriteria): Promise<number>;
isExist(id: number): Promise<boolean>;
checkExist(id: number): Promise<void>;
softDeleteById(id: number): Promise<undefined>;
bulkSoftDelete(criteria: ICriteria): Promise<undefined>;
}
*/
type ICriteria struct {
	Select []string
	filters []IFilter
	sort ISort
	offset int
	limit int
	includes []IInclude
}

type IInclude struct {
	field string
	Select []string
	includes []IInclude
	filters []IFilter
}

type IFilter struct {
	code string
	operator string
	value string
}

type SortDirection int

const (
	ASC SortDirection = iota
	DESC
)

func (s SortDirection) String() string {
	return [...]string{"asc", "desc"}[s]
}



type ISort struct {
	column string
	direction SortDirection
}
