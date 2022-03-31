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