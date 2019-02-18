-- 第二高的薪水
-- 倒数第 2 即按倒序排取第 2 个
-- 内部子查询返回空结果，外部 AS 取结果值为 NULL
select 
    (select distinct Salary from Employee 
            order by Salary desc limit 1,1) 
as SecondHighestSalary;