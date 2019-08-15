<script>
  import MeetUpItem from "./MeetUpItem.svelte";
  import MeetupFilter from "./MeetupFilter.svelte";
  import Button from "../UI/Button.svelte";
  import { createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";
  import { flip } from "svelte/animate";

  const dispatch = createEventDispatcher();

  export let meetups;

  let filter = false;

  $: filteredMeetups = filter
    ? filter === 1
      ? meetups.filter(m => m.isLiked)
      : meetups.filter(m => !m.isLiked)
    : meetups;

  const setFilter = e => {
    filter = e.detail;
  };
</script>

<style>
  #meetups {
    width: 100%;
    display: grid;
    grid-template-columns: 1fr;
    grid-gap: 1rem;
  }

  #meetup-controls {
    margin: 1rem;
    display: flex;
    justify-content: space-between;
  }

  @media (min-width: 768px) {
    #meetups {
      grid-template-columns: repeat(3, 1fr);
    }
  }
</style>

<section id="meetup-controls">
  <MeetupFilter on:select={setFilter} />
  <Button on:click={() => dispatch('add')}>New Meetup</Button>

</section>
<section id="meetups">


  {#each filteredMeetups as meetup (meetup.ID)}
    <div transition:scale animate:flip={{duration:300}}>
      <MeetUpItem
        id={meetup.ID}
        title={meetup.title}
        subtitle={meetup.subtitle}
        description={meetup.description}
        imageUrl={meetup.imageUrl}
        contactEmail={meetup.contactEmail}
        isLiked={meetup.isLiked}
        address={meetup.address}
        on:edit
        on:showdetails />
    </div>
  {/each}
</section>
