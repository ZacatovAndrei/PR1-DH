class Waiter
{
    public Waiter(int id) => this.id = id;
    private Order? currentOrder;
    public int id { get; private set; }

    public void takeOrder(Table[] tables)
    {
        if (this.currentOrder is not null) return;
        foreach (var table in tables)
        {
            if (table.isEmpty && table.orderState == OrderStates.READY)
            {
                currentOrder=new Order(table.tableOrder)
                Console.WriteLine($"order received from table {table.id}");
            }
            else Thread.Sleep(250);
        }
    }
}