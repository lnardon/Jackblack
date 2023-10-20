<script>
let rooms = [{id: 1, card: "2_of_clubs"}, {id: 2, card: "9_of_spades"}]

</script>

<div>
    {#each rooms as room}
        <div>
            <img src={"assets/cards/"+room.card+".png"} alt="Card: ">
        </div>
    {/each}
</div>

<style>
    div {
        display: flex;
        justify-content: space-around;
        flex-wrap: wrap;
        gap: 2rem;
    }

    img {
        width: 200px;
        height: 300px;
        transition: all 0.9s cubic-bezier(0.075, 0.82, 0.165, 1);
        animation: spin 1.1s cubic-bezier(0.175, 0.885, 0.32, 1.275);
        cursor: pointer;
    }

    img:hover {
        transform: scale(1.05);
    }

    @keyframes spin {
        0% { transform: rotateY(145deg); 
            opacity: 0;
        }
        100% { transform: rotateY(0deg); 
            opacity: 1;
        }
    }
</style>