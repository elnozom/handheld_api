USE mcstree

GO
ALTER PROC CashtrayClose(
    @Serial int,
    @Exceed real,
    @Shortage real,
    @Amount real
)
AS
BEGIN 
    UPDATE CashTry SET Exceed = @Exceed ,Shortage = @Shortage, CasherMoney = @Amount , CloseData = GETDATE()  WHERE "Serial" = @Serial
    SELECT @Serial  "serial";
END
