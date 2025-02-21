Golang sample_scaffold_working_structure

<pre>
% go run main.go 
Original string: you should be able to see me
Hex representation: 796f752073686f756c642062652061626c6520746f20736565206d65
Before deletion, raw memory read: you should be able to see me
After deletion, raw memory read: /* the garbarge collection is too fast for us to catch the memory remnants */
After zeroization, reading memory: /* the garbarge collection is too fast for us to catch the memory pointer */

</pre>

