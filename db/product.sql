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
ALTER PROCEDURE StkMs01FindByCode (@BCode  NVARCHAR(20)= null ,@StoreCode int, @Name NVARCHAR(50) = '' )
AS
    DECLARE @m_Code  FLOAT
	DECLARE @GC INT
	DECLARE @IC INT
	DECLARE @ExpireCode NVARCHAR(20)
	DECLARE @Code NVARCHAR(20) 
	DECLARE @ExpireD NVARCHAR(20)
    
    -- #check if barcode is smaller than or equal 6 chars to extract the group code and item code
    -- #this because that any barcode smaller than or equal 6 chars will be local barcode not international 
    IF LEN(@BCode) <= 6 AND @BCode IS NOT NULL
        BEGIN
            SET @m_Code = CAST( @BCode AS INT)
            SET  @GC = right(format (@m_Code ,'000000') ,2)
            SET  @IC = Left(format (@m_Code ,'000000') ,4)
        END
    ELSE
        BEGIN
            IF Left(@BCode, 1) = 2 And Len(@BCode) = 11 
                BEGIN
                    SET @ExpireCode =  SUBSTRing (@BCode,2,10)
                    SET  @Code = Left(@ExpireCode, 6)
                    SET  @ExpireD = Right(@ExpireCode, 4)
                    SET @m_Code = CAST( @code AS INT)
                    SET  @GC = right(format (@m_Code ,'000000') ,2)
                    SET  @IC = Left(format (@m_Code ,'000000') ,4)
                END
        END 

    
    --# select data depending on our vars
    select  ItemCode ,
            GroupCode ,
            StkMs01.BarCode ,
            POSName ,
            ItemTypeID ,
            MinorPerMajor ,
            StkMs01.AccountSerial ,
            ActiveItem ,
            ItemHaveSerial ,
            MasterItem ,
            ItemHaveAntherUint,
            StoreCode,
            StkMs02.LastBuyPrice PriceFinal, 
            ISNULL(ItemAccount.LastBuyPrice , 0) PriceBefore, 
            ISNULL(ItemAccount.Disc1 , 0) Disc1, 
            ISNULL(ItemAccount.Disc2 , 0) Disc2, 
            ISNULL(ItemAccount.Tax1 , 0) Tax1, 
            POSTP,
            POSPP,
            ISNULL( AccMs01.AccountCode , 0) AccountCode,
            ISNULL( AccMs01.AccountName  ,'') AccountName,
            Ratio1,
            Ratio2
    FROM StkMs01
    JOIN StkMs02 ON StkMs01.Serial = StkMs02 .ItemSerial 
    LEFT OUTER JOIN ItemAccount ON StkMs01.Serial = ItemAccount.AccountSerial 
    LEFT OUTER JOIN AccMs01 ON AccMs01.Serial = StkMs01.AccountSerial 
    LEFT OUTER JOIN BarCode ON StkMs01.Serial = BarCode.ItemSerial 
    WHERE StkMs02.StoreCode = @StoreCode 
    AND
    ItemName LIKE ( '%' + CASE WHEN @Name != '' THEN  @Name ELSE ItemName END +'%') 
    AND
    ItemCode = CASE WHEN @IC IS NULL THEN ItemCode ELSE @IC END  
    AND  
    StkMs01.GroupCode = CASE WHEN @IC IS NULL THEN StkMs01.GroupCode ELSE @GC END
    AND
    (
        StkMs01.BarCode = CASE WHEN  @IC IS NULL THEN @BCode   ELSE  StkMs01.BarCode END OR
        BarCode.ItemBarCode = CASE WHEN  @IC IS NULL THEN @BCode  ELSE BarCode.ItemBarCode END 
    )

   
	
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

