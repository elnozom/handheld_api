USE MCSTREE

GO
ALTER PROC EmployeeLogin(
    @EmpCode INT,
    @EmpPassword VARCHAR(200)
)
AS
BEGIN 
    SELECT EmpCode FROM Employee WHERE EmpCode = @EmpCode AND EmpPassword = @EmpPassword
END
