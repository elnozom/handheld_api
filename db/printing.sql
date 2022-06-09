USE MCSTREE

GO
-- list order and items by item serial
-- i refers to item , oi refers to orderItem
ALTER PROCEDURE [dbo].[StkTr01PrintItemsBySerial] (@Serial int)
AS
BEGIN
SELECT 
 CAST( DocDate AS DATE  ) Docdate ,CAST (DocDate as time ) DocTime,    AccountName , DocNo,ItemName ,EmpName,
  sum(Qnt)TQnt ,StkTr02.MinorPerMajor ,Price ,Sum(Qnt * Price ) totalPrice
			 from StkTr01
			 inner join Employee on  StkTr01.CasherCode = Employee.EmpCode 
			 inner join StkTr02  on StkTr01.Serial = StkTr02.HeadSerial 
			 inner join AccMs01 on StkTr01 .AccountSerial = AccMs01.Serial 
			 inner join StkMs01 on StkTr02.ItemSerial = StkMs01.Serial 
	WHERE StkTr01 .Serial  = @Serial
	GROUP BY ItemSerial , DocNo, DocDate, DocTime,  AccountName , OrderNo,  BonNo,
  TableNo, Price ,ItemName ,StkTr02.MinorPerMajor
  , EmpName
	
END

