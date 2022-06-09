USE MCSTREE



GO
CREATE PROCEDURE StkTr01ListItems (@Serial INT)
AS
BEGIN
    DECLARE @totalCash FLOAT
	SET @totalCash = (SELECT SUM(Qnt * Price / MinorPerMajor) FROM StkTr02 WHERE HeadSerial = @Serial)
    SELECT 
        d.Serial,
        d.Qnt,
        d.Price,
        i.BarCode,
		@totalCash,
        i.ItemName,
        i.MinorPerMajor,
        (d.Price * (d.Qnt / d.MinorPerMajor)),
        i.ByWeight1
    FROM StkTr02 d JOIN StkMs01 i ON d.ItemSerial = i.Serial WHERE HeadSerial = @Serial

END


GO
CREATE PROCEDURE StkTr02Delete (@Serial INT)
AS
BEGIN
    DELETE FROM StkTr02 WHERE Serial = @Serial
    SELECT 1 deleted
END

GO
ALTER PROCEDURE StkTr01List (@TransSerial INT , @StoreCode INT , @ComputerName VARCHAR(200))
AS
BEGIN
    SELECT o.Serial, o.StoreCode , o.DocNo , o.AccountSerial , o.TransSerial ,(SELECT SUM(Qnt * Price) FROM StkTr02 WHERE HeadSerial = o.Serial) , a.AccountName , a.AccountCode FROM StkTr01 o JOIN AccMs01 a ON o.AccountSerial = a.Serial WHERE o.StoreCode = @StoreCode AND o.TransSerial = @TransSerial AND o.ComputerName = @ComputerName
END


GO
ALTER PROCEDURE StkTr01Insert(
    @AccountSerial INT,
    @EmpCode INT,
    @StoreCode INT,
    @StoreCode2 INT = 0,
    @ComputerName VARCHAR(100),
    @HeadSerial INT,
    @TransSerial INT,
    @ItemSerial INT,
    @Qnt FLOAT,
    @Price FLOAT,
    @Tax FLOAT = 0,
    @MinorPerMajor INT
)
AS
    BEGIN
        IF @HeadSerial = 0 
        BEGIN
            INSERT INTO StkTr01
                (
                    DocNo,
                    CasherCode,
                    AccountSerial,
                    StoreCode,
                    StoreCode2,
                    ComputerName,
                    TransSerial
                ) VALUES (
                    dbo.StkTr01GetMaxDocNo(@StoreCode , @TransSerial),
                    @EmpCode,
                    @AccountSerial,
                    @StoreCode,
                    @StoreCode2,
                    @ComputerName,
                    @TransSerial
                )
            
           SET @HeadSerial = (SELECT SCOPE_IDENTITY())
        END
        INSERT INTO StkTr02
            (
                HeadSerial ,
                ItemSerial ,
                Qnt ,
                Price ,
                Tax ,
                MinorPerMajor
            ) VALUES (
                @HeadSerial,
                @ItemSerial,
                @Qnt,
                @Price,
                @Tax,
                @MinorPerMajor
            )
            SELECT @HeadSerial headSerial
    END


GO
ALTER FUNCTION StkTr01GetMaxDocNo( @StoreCode INT) returns nvarchar(10)  
AS  
BEGIN 

    DECLARE @DocNo INT
    DECLARE @StoreLetter CHAR
    SET @StoreLetter = (SELECT FixedChar FROM StoreCode WHERE StoreCode = 1)
    SET @DocNo = 
		(SELECT TOP(1) CONVERT(INT, SUBSTRING(DocNo, 2, LEN(DocNo) - 1))  
			FROM StkTr01 
			WHERE DocNo LIKE '%'+ @StoreLetter +'%' 
				AND DocNo NOT like '%-%' AND TransSerial = 19  
				AND TransIndex = 1 
			ORDER BY CONVERT(INT, SUBSTRING(DocNo, 2, LEN(DocNo) - 1))Desc)
   IF @DocNo IS NULL 
		SET @DocNo = 0 

    RETURN CONCAT(@StoreLetter , @DocNo + 1)
END  


