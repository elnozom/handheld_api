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
ALTER PROCEDURE StkMs01FindByCode (@BCode  NVARCHAR(20)= null ,@StoreCode int, @Name NVARCHAR(50) = null )
AS

    DECLARE @m_Code  FLOAT
	DECLARE @GC INT
	DECLARE @IC INT
	DECLARE @ExpireCode NVARCHAR(20)
	DECLARE @Code NVARCHAR(20) 
	DECLARE @ExpireD NVARCHAR(20)



if @Name is not null
	BEGIN 
		select  ItemCode ,
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
		from StkMs01
		inner join StkMs02 on StkMs01.Serial = StkMs02 .ItemSerial 
		where ItemName like ( '%' + @Name +'%') and StkMs02.StoreCode = @StoreCode
	end
else if Len(@BCode) <= 6 and @BCode is not null
 begin
 SET @m_Code = CAST( @BCode AS INT)
			set  @GC = right(format (@m_Code ,'000000') ,2)
		    set  @IC = Left(format (@m_Code ,'000000') ,4)
			select  ItemCode ,
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
			from StkMs01
			inner join StkMs02 on StkMs01.Serial = StkMs02 .ItemSerial 
			where (ItemCode = @IC and StkMs01.GroupCode = @GC) and StkMs02.StoreCode = @StoreCode
			end 
else
begin
IF Left(@BCode, 1) = 2 And Len(@BCode) = 11 
begin
				set @ExpireCode =  SUBSTRing (@BCode,2,10)
		
                SET  @Code = Left(@ExpireCode, 6)
                SET  @ExpireD = Right(@ExpireCode, 4)
		    SET @m_Code = CAST( @code AS INT)
		    set  @GC = right(format (@m_Code ,'000000') ,2)
		    set  @IC = Left(format (@m_Code ,'000000') ,4)
            select  
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
			from StkMs01
			inner join StkMs02 on StkMs01.Serial = StkMs02 .ItemSerial 
			where (ItemCode = @IC and StkMs01.GroupCode = @GC) and StkMs02.StoreCode = @StoreCode 
	end 		
			 

ELSE	

	begin
			select  ItemCode ,
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
			from StkMs01
			INNER JOIN StkMs02 on StkMs01 .Serial = StkMs02.ItemSerial 
			left outer join BarCode on StkMs01.Serial = BarCode.ItemSerial 
			where 
			(BarCode = @BCode or BarCode.ItemBarCode = @BCode) and StkMs02.StoreCode = @StoreCode

    end 
 
 end
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

