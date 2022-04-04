USE mcstree

GO
ALTER PROC StkMs01CreateInitial(@ItemCode INT , @GroupCode INT , @BarCode VARCHAR(20),@Name VARCHAR(200),@MinorPerMajor INT)
AS
BEGIN 
INSERT INTO StkMs01 (
    ItemCode , GroupCode , BarCode , POSName , ItemName ,MinorPerMajor
) VALUES (
    @ItemCode , @GroupCode , @BarCode , @Name , @Name ,@MinorPerMajor
)
SELECT SCOPE_IDENTITY() serial
END



GO
ALTER PROC StkMs01MacItemCodeByGroup(@GroupCode TINYINT)
AS
BEGIN 
SELECT (MAX(ItemCode) + 1 ) maxCode , g,GroupName FROM StkMs01 i JOIN GroupCode g ON g.GroupCode = i.GroupCode WHERE GroupCode = @GroupCode 
SELECT SCOPE_IDENTITY() serial 
END


GO
CREATE PROC ItemTypeByGroup(@GroupCode TINYINT)
AS
BEGIN 
SELECT ItemTypeID , ItemTypeName  FROM ItemType it WHERE GroupCode  = @GroupCode
END





    GO
    ALTER PROC StkMs01FindByCode(@ItemCode SMALLINT , @GroupCode TINYINT)
    AS
    BEGIN 
        SELECT 
        TOP(1)
            ItemCode ,
            GroupCode ,
            BarCode ,
            POSName ,
            ItemTypeID ,
            MinorPerMajor ,
            AccountSerial ,
            ActiveItem ,
            ItemHaveSerial ,
            MasterItem ,
            ItemHaveAntherUint,
            StoreCode,
            LastBuyPrice,
            POSTP,
            POSPP,
            Ratio1,
            Ratio2
            FROM StkMs01 
            JOIN StkMs02 ON StkMs01.Serial = StkMs02.ItemSerial
            WHERE GroupCode = @GroupCode AND ItemCode = @ItemCode
    END

