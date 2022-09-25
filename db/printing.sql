
GO
ALTER PROCEDURE [dbo].[StkTr01PrintItemsBySerial] (@Serial int)
AS
BEGIN
	SELECT i.DocDate , ISNULL(i.Discount ,0 ), a.AccountName , ISNULL(a.Address , '') ,i.DocNo ,e.EmpName,s.StoreName,
			FROM StkTr01 i
			JOIN Employee e on  i.CasherCode = e.EmpCode
			JOIN StoreCode s ON  i.StoreCode = s.StoreCode
			JOIN AccMs01 a on i .AccountSerial = a.Serial 
	WHERE i.Serial  = @Serial


	SELECT  oi.ItemName ,
		CASE WHEN oi.MinorPerMajor = 1000 THEN Qnt else   FLOOR(Qnt / oi.MinorPerMajor ) END  wholeQnt ,
		CASE WHEN oi.MinorPerMajor = 1000 THEN 0  else (CAST(Qnt AS INT) %  CAST(oi.MinorPerMajor AS INT)) END partQnt,
		oi.MinorPerMajor,
		Price ,
		Sum (CASE WHEN oi.MinorPerMajor = 1000  THEN Qnt * Price  ELSE (Qnt * Price / oi.MinorPerMajor ) END) totalPrice,
		i.BarCode 
			from StkTr02 oi
			inner join StkMs01 i on oi.ItemSerial = i.Serial 
	WHERE oi.HeadSerial  = @Serial
END



-- USE [HalalDubai]
-- GO
-- ALTER PROCEDURE [dbo].[StkTr01PrintItemsBySerial] (@Serial int)
-- AS
-- BEGIN
-- SELECT DocDate , ISNULL(StkTr01.Discount ,0 ), AccountName , ISNULL(AccMs01.Address , '') ,
-- DocNo,ItemName ,EmpName,StoreCode.StoreName,
--   case when StkTr02.MinorPerMajor = 1000 then Qnt else   FLOOR(Qnt / StkTr02.MinorPerMajor ) end  wholeQnt ,
--    case when StkTr02.MinorPerMajor = 1000 then 0  else (CAST(Qnt AS INT) %  CAST(StkTr02.MinorPerMajor AS INT)) end partQnt   ,
--   StkTr02.MinorPerMajor ,Price ,
--   Sum (case when  StkTr02.MinorPerMajor = 1000  then Qnt * Price   else (Qnt * Price / StkTr02.MinorPerMajor ) end) totalPrice,
--   StkMs01.BarCode 
-- 			 from StkTr01
-- 			 inner join Employee on  StkTr01.CasherCode = Employee.EmpCode
-- 			 inner join StoreCode ON  StkTr01.StoreCode = StoreCode.StoreCode
-- 			 inner join StkTr02  on StkTr01.Serial = StkTr02.HeadSerial 
-- 			 inner join AccMs01 on StkTr01 .AccountSerial = AccMs01.Serial 
-- 			 inner join StkMs01 on StkTr02.ItemSerial = StkMs01.Serial 
-- 	WHERE StkTr01 .Serial  = @Serial
-- 	GROUP BY AccMs01.Address,StkTr01.Discount , ItemSerial ,StoreCode.StoreName,Qnt, DocNo, DocDate,  AccountName , OrderNo,  BonNo,
--   TableNo, Price ,ItemName ,StkTr02.MinorPerMajor,StkMs01.BarCode 
--   , EmpName
	
-- END
