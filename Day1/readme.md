// 1) What if the number of seats is unknown?
// For first answer i am thinking to use slice . Because slice is dyanmic so it can grow . But tradeoff is growth algorithm cause reallocation which cause gc pressure . Now so first we try to fix correct len and capacity for ideal case of flight so that no of reallocations should be less.


//2) How to handle multiple bookings at the same time?
// Here we can use mutex to avoid race condition either mutex between whole book ticket operations (means consider that as shared memory)  or we can do different seat wise mutex but this looks much complex too manage , as all seats has mutex . 


// 3) Would a linked list be better in some cases?
// In this particular problem linked list is not looking optimal approach as here we need to book ticket based per seat so in array we directly use index O(1) operation . But in linked list we need to iterate through it which is O(n) not optimal . Here we also don't need middle insertion or deletion as per our usecase . So final judgement should be fixed arrays or slices is more beneficial .
