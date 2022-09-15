enum OrderStates
{
    NONE, READY, WAITING
};

class Table
{
    public int id { get; private set; }
    public bool isEmpty { get; private set; }
    public OrderStates orderState { get; private set; }
    public Order? tableOrder { get; private set;}

    public Table(int id) => (this.id, this.isEmpty, this.orderState) = (id, true, OrderStates.NONE);

    public void occupy()
    {
        this.isEmpty = false;
    }
    public void free()
    {
        this.isEmpty = true;
    }
    public void generateOrder()
    {
        //sanity check
        if (this.tableOrder is not null)
        {
            Console.WriteLine($"Table#{this.id}:Order already exists!");
            return;
        }
        //intialising order info

        var rand = new Random();
        this.tableOrder = new Order();
        tableOrder.table_id = this.id;
        tableOrder.order_id = Globals.FreeOrderNumber++;

        //generating the order contents

        tableOrder.items.Add(rand.Next(1, 13));

        int randint = rand.Next(0, 13);
        while (randint != 0)
        {
            tableOrder.items.Add(randint);
            randint = rand.Next(1, 13);
        }

    }
}