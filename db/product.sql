USE mcstree
GO
ALTER PROC StkMs01CreateInitial
(@ItemCode INT ,
@GroupCode INT , 
@BarCode VARCHAR(20),
@Name VARCHAR(200),
@MinorPerMajor INT,
@AccountSerial int,
@ActiveItem bit,
@ItemTypeID int ,
@ItemHaveSerial bit,
@MasterItem bit,
@ItemHaveAntherUint bit,
@StoreCode int ,
@LastBuyPrice real,
@POSTP real,
@POSPP real ,
@Ratio1  real,
@Ratio2 real)
AS
BEGIN

declare @ItemSerial int  
	 SET @ItemSerial = (SELECT Serial FROM StkMs01 WHERE ItemCode = @ItemCode AND GroupCode = @GroupCode)
	 IF @ItemSerial IS NOT NULL
		BEGIN
		  UPDATE StkMs01 
			SET BarCode  =@BarCode ,
			 POSName  =@Name ,
			 ItemName  =@Name ,
			 ItemTypeID =@ItemTypeID ,
			 MinorPerMajor =@MinorPerMajor ,
			 AccountSerial =@AccountSerial ,
			 ActiveItem =@ActiveItem ,
			 ItemHaveSerial =@ItemHaveSerial ,
			 MasterItem =@MasterItem ,
			 ItemHaveAntherUint =@ItemHaveAntherUint 

		WHERE 
		ItemCode = @ItemCode AND GroupCode = @GroupCode




		UPDATE StkMs02
					SET ItemSerial = @ItemSerial,
					StoreCode = @StoreCode,
					LastBuyPrice = @LastBuyPrice,
					AvrPrice = @LastBuyPrice,
					POSTP = @POSTP,
					POSPP = @POSPP,
					Ratio1 = @Ratio1,
					Ratio2 = @Ratio2,
					Percnt1 = 1,
					Percnt2 =1 

		WHERE 
		ItemSerial = @ItemSerial

        SELECT @ItemSerial serial
		RETURN 
	END

INSERT INTO StkMs01 (
    ItemCode , GroupCode , BarCode , POSName , ItemName ,ItemTypeID,MinorPerMajor,AccountSerial,
	ActiveItem,ItemHaveSerial,MasterItem,ItemHaveAntherUint
)
 VALUES 
(
    @ItemCode , @GroupCode , @BarCode , @Name , @Name ,@ItemTypeID,@MinorPerMajor,
	@AccountSerial,@ActiveItem ,@ItemHaveSerial,@MasterItem,@ItemHaveAntherUint
)

set @ItemSerial = SCOPE_IDENTITY()   


INSERT INTO StkMs02
             (ItemSerial, StoreCode,  LastBuyPrice, AvrPrice,   POSTP, POSPP, Ratio1, Ratio2, Percnt1, Percnt2)
VALUES   (@ItemSerial,@StoreCode,@LastBuyPrice,@LastBuyPrice,@POSTP,@POSPP,@Ratio1,@Ratio2,1,1)


SELECT @ItemSerial serial
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

