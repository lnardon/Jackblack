<script>
  import svelteLogo from '../public/assets/svelte.svg'
  import Counter from './lib/Counter.svelte'
  import Cards from './lib/Cards.svelte';

  let currentRoom = 0;
  let roomOwner = true;
  let cards =[]

  function setCurrentRoom(roomId){
    currentRoom = roomId;
  }

  const conn = new WebSocket("ws://localhost:8080/ws");
  
  conn.onclose = (evt) => {
    console.log('CONNECTION CLOSED', evt)
  }

  conn.onmessage = (evt) => {
    console.log('MESSAGE RECEIVED', evt.data);

    if(evt.data.startsWith("room")){
      setCurrentRoom(evt.data.split(":")[1])
    } else {
      cards = [...cards, ...evt.data.split("\n")]
      console.log(cards)
    }
    
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

  function startGame(){
    conn.send("start_game")
  }

</script>

<main>
  <div>
    <a href="https://svelte.dev" target="_blank" rel="noreferrer">
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
  </div>
  <h1>Jackblack <br/>(Online Blackjack)</h1>

  <div class="card">
    {#if currentRoom == 0}
    <Counter createRoom={createRoom} joinRoom={joinRoom} getAllRooms={getAllRooms}/>
    {:else}
      <div>
        <h2>Room ID: {currentRoom}</h2>
        <div>
          <h3>Players</h3>
          <ul>

          </ul>
      </div>
            {#if roomOwner}
      <button on:click={startGame}>
        Start Game
      </button>
      <Cards cards={cards}/>
      {/if}
  </div>
  {/if}
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 700ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }

  button {
    margin-bottom: 2rem;
  }
</style>
