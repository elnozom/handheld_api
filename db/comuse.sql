USE MCSTREE


-- ALTER TABLE ComUse ADD Imei  VARCHAR(200) UNIQUE
-- ALTER TABLE ComUse ADD CONSTRAINT U_Name UNIQUE(ComName)
-- ALTER TABLE ComUse ADD CONSTRAINT U_Imei UNIQUE(Imei)
GO
ALTER PROC ComUseCreate(
    @ComName VARCHAR(200),
    @Imei VARCHAR(200)
)
AS
BEGIN 
    DECLARE @Capital Char
	DECLARE @Store Char
	SELECT @Capital = CHAR(ASCII(MAX(Capital)) + 1)  FROM ComUse 
   INSERT INTO ComUse (
        ComName,
        Capital,
        Imei
    ) VALUES (
        @ComName,
        @Capital,
        @Imei
    )
    SELECT Imei , ComName , Capital FROM ComUse WHERE Serial1 =SCOPE_IDENTITY()

END



GO
ALTER PROC ComUseFind(@Imei VARCHAR(200) = NULL ,@ComName VARCHAR(200) = NULL )
AS
BEGIN 
   SELECT Imei , ComName , Capital FROM ComUse WHERE Imei = CASE WHEN  @Imei IS NULL THEN Imei ELSE  @Imei END AND ComName = CASE WHEN  @ComName IS NULL THEN ComName ELSE  @ComName END
END
