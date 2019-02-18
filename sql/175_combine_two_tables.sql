-- 组合两张表
-- 相比 where 的真过滤，on 是假过滤，会保留一侧的无数据字段为 Null
select FirstName, LastName, City, State
	from Person as p left join Address as a 
					 on p.PersonId = a.PersonId;