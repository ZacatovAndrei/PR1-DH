/*
*
* The whole info about deserialising json has been taken from here
* https://docs.microsoft.com/en-us/dotnet/standard/serialization/system-text-json-how-to?pivots=dotnet-6-0
*
*/
class Order
{
    public int order_id;
    public int table_id;
    public int waiter_id;
    public List<int> items;
    public int priority;
    public int max_wait;
    public int pickup_time;
    public int cooking_time;
    public List<Dictionary<String, int>>? cooking_details;
    public Order()
    {
        items = new List<int>();
    }
    public Order(Order a)
    {
        order_id = a.order_id;
        table_id = a.table_id;
        waiter_id = a.waiter_id;
        memberw


    }
}
