<script>
  export let setCurrentRoom;
  const conn = new WebSocket("ws://localhost:8080/ws");
  
  conn.onclose = (evt) => {
    console.log('CONNECTION CLOSED', evt)
  }

  conn.onmessage = (evt) => {
    console.log('MESSAGE RECEIVED', evt.data);
    setCurrentRoom(evt.data)
  }

  function createRoom(){
    conn.send("create_room")
  }

  function joinRoom(){
    let roomId = prompt("Enter room id")
     conn.send("join_room:"+roomId)
  }

  function getAllRooms(){
    conn.send("get_all_rooms")
  }

</script>

<div>
  <button on:click={joinRoom}>
    Join Room
  </button>
  <button on:click={createRoom}>
    Create room
  </button>
  <button on:click={getAllRooms}>
    Get rooms
  </button>
</div>
