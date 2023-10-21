<script>
  export let createRoom;
  export let joinRoom;
  export let getAllRooms;

</script>

<div class="container">
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

<style>
  .container{
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    gap: 1rem;
  }
</style>