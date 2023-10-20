<script>
  const conn = new WebSocket("ws://localhost:8080/ws");
  
  conn.onclose = (evt) => {
    console.log('CONNECTION CLOSED', evt)
  }

  conn.onmessage = (evt) => {
    console.log('MESSAGE RECEIVED', evt.data);
  }

  function createRoom(){
    conn.send("create_room")
  }

  function joinRoom(){
    let roomId = prompt("Enter room id")
     conn.send(`join_room:${roomId}`)
  }

</script>

<div>
  <button on:click={joinRoom}>
    Join Room
  </button>
  <button on:click={createRoom}>
    Create room
  </button>
</div>
