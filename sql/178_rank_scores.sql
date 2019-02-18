-- 分数排名
-- 排名即，去重后有多少 rows 比我的值还大
select Score, (select count(distinct Score) from Scores where Score >= s.Score) as Rank
from Scores as s 
order by Score desc