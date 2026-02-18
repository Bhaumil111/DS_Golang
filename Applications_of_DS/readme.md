# Data Structures — Practical Usage Guide
 
---

## 1. Arrays (Fixed Size)

**When size is known, memory contiguous, fastest access needed**

- Fixed-size buffers (network packets, audio/video frames)
- Static lookup tables (ASCII map, config flags)
- Matrix / grid storage (images, game boards)
- CPU core / worker slot mapping
- Embedded systems / memory-constrained programs
- Ring buffers (circular array implementation)
- Cache-friendly high-performance loops

---

## 2. Dynamic Arrays (Resizable Arrays / Slices / Vectors)

**When size grows/shrinks dynamically**

- Dynamic collections (user lists, results, logs)
- Batch processing (DB batch insert, message batching)
- Buffers for IO / file / network streaming
- Collecting results from parallel workers
- Dynamic stacks and queues
- In-memory tables (rows stored in array-like structure)
- JSON / request / response aggregation
- Game objects / entity lists

---

## 3. Singly Linked List

**When frequent insert/delete and sequential traversal**

- Dynamic memory allocation required
- Implementing queues (simple FIFO)
- Memory-efficient large sequential data
- Streaming pipelines
- Chaining in hash-table collision handling
- Undo-like forward-only logs
- Task/job chaining
- Adjacency list in graph (simple representation)

---

## 4. Doubly Linked List

**When forward + backward traversal required**

- Dynamic allocation necessary
- Prev/Next traversal scenarios
- Implementing dynamic deque (double-ended queue)
- Implementing LRU / MRU caches
- Maintaining histories (undo/redo)
- Navigation systems (browser back/forward)
- Playlist / ordered iteration with removal
- OS page replacement / buffer cache
- Ordered eviction policies

---

## 5. Hash Maps (Hash Tables / Dictionaries)

**Fast key → value lookup (O(1) average)**

- In-memory databases / caches
- Session/token storage
- Indexing (userID → data)
- Deduplication / visited tracking
- Frequency counting / aggregation
- Routing / lookup tables
- Symbol tables (compilers/interpreters)
- Memoization / dynamic programming cache
- Configuration / environment storage
- Graph adjacency map
- Real-time analytics counters

---

## 6. Queues (FIFO)

**First-in-first-out processing**

- Task scheduling / job queues
- Producer–consumer pipelines
- Request buffering (servers, message brokers)
- BFS graph traversal
- Streaming systems
- Event processing
- OS scheduling
- Rate limiting / token bucket
- Print queues / batch jobs

Variants:
- Circular queue → fixed buffers
- Priority queue → scheduler
- Concurrent queue → worker pool

---

## 7. Stacks (LIFO)

**Last-in-first-out**

- Function call stack (runtime)
- Undo/Redo systems
- Expression parsing / evaluation
- DFS graph traversal
- Backtracking algorithms
- Syntax parsing (compilers)
- Balanced parentheses / bracket checking
- Memory stack allocation
- State rollback systems

---

## 8. Trees (General Hierarchical Structure)

**Hierarchical / ordered data**

- File systems / directory structure
- XML/HTML/JSON DOM
- Database indexing (B/B+ trees)
- Expression trees (compilers)
- Organization hierarchy
- Game scene graph
- Routing / decision trees
- Interval trees (range queries)
- Segment trees (range computation)
- Binary search tree (ordered data)

---

## 9. Heaps (Priority Queues)

**Efficient min/max extraction**

- Task scheduling by priority
- Job schedulers
- Event simulation (next event)
- Dijkstra / shortest path
- Load balancing
- Top-K problems
- Streaming median
- Memory management
- Timer / delayed execution queues
- AI / pathfinding

---

## 10. Graphs

**Relationships / networks**

- Social networks (followers, friends)
- Road maps / navigation
- Recommendation systems
- Dependency graphs (build systems, package managers)
- Distributed systems / service mesh
- Web crawling
- Network routing
- State machines
- Fraud detection / anomaly detection
- Knowledge graphs
- Resource allocation graphs (deadlock detection)

---

## 11. Tries (Prefix Trees)

**Fast prefix / string search**

- Autocomplete / search suggestions
- Spell checker
- Dictionary / word lookup
- IP routing (longest prefix match)
- Text search engines
- Contact search (phone, email)
- URL routing
- Token matching / lexical analysis
- DNA / genome sequence matching

---

## 12. Deque (Double Ended Queue)

**Insert/remove both ends**

- Sliding window algorithms
- Task scheduling
- Cache eviction
- Monotonic queue (range max/min)
- BFS with bidirectional traversal
- Undo/Redo buffer
- Real-time stream window

---

## 13. Set (Hash Set / Tree Set)

**Unique elements only**

- Deduplication
- Membership testing
- Graph visited tracking
- Unique user/session tracking
- Recommendation filtering
- Permission systems
- Fast lookup existence

---

## 14. Bitset / Bitmap

**Memory-efficient boolean storage**

- Feature flags
- Permission bits
- Bloom filters
- Fast membership test
- Large-scale analytics
- Compressed state storage

---

## 15. Bloom Filter (Probabilistic DS)

**Fast approximate membership**

- Cache existence check
- Duplicate filtering
- Database / disk avoidance
- Distributed systems
- Web crawler visited check
- Big-data pipelines

---

## 16. Ring Buffer (Circular Buffer)

**Fixed-size streaming**

- Logging systems
- Telemetry pipelines
- Audio/video streaming
- High-performance queues
- Lock-free communication
- Real-time systems

---

## Quick Pattern Summary (How Engineers Think)

| Need                   | Data Structure     |
|------------------------|-------------------|
| Fast index access      | Array             |
| Dynamic list           | Dynamic array     |
| Frequent insert/delete | Linked list       |
| Forward + backward     | Doubly linked list|
| Fast lookup by key     | Hash map          |
| FIFO processing        | Queue             |
| LIFO processing        | Stack             |
| Hierarchical           | Tree              |
| Priority scheduling    | Heap              |
| Relationships/network  | Graph             |
| Prefix search          | Trie              |
| Sliding window         | Deque             |
| Unique items           | Set               |
| Memory efficient flags | Bitset            |
| Approx membership      | Bloom filter      |
| Streaming buffer       | Ring buffer       |