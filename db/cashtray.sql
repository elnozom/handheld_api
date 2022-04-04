USE mcstree

GO
CREATE PROC CashtrayClose(
    @Serial int,
    @Exceed real,
    @Shortage real,
    @Amount real
)
AS
BEGIN 
    UPDATE CashTry SET Exceed = @Exceed ,Shortage = @Shortage, CasherMoney = @Amount WHERE "Serial" = @Serial
    SELECT @Serial  "serial";
END
