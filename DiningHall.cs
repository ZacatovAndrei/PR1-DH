class DiningHall
{
    private int waiterNum;
    private Waiter[] waiterArr;
    private int tableNum;
    private Table[] tableArr;
    public DiningHall(int tables, int waiters)
    {
        Console.WriteLine($"Dinnerhall: initialising a Dining Hall with {tables} tables and {waiters} waiters");

        (tableNum, waiterNum) = (tables, waiters);
        // generating array of tables, each with their own ID
        tableArr = new Table[tableNum];
        for (int i = 0; i < tableNum; i++)
        {
            tableArr[i] = new Table(i);
        }
        // generating array of waiters, each with their own ID
        waiterArr = new Waiter[waiterNum];
        for (int i = 0; i < waiterNum; i++)
        {
            waiterArr[i] = new Waiter(i);
        }

    }

}