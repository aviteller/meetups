<script>
  import Header from "./UI/Header.svelte";
  import MeetUpGrid from "./Meetups/MeetUpGrid.svelte";
  import EditMeetup from "./Meetups/EditMeetup.svelte";
  import MeetupDetail from "./Meetups/MeetupDetail.svelte";
  import TextInput from "./UI/TextInput.svelte";
  import Button from "./UI/Button.svelte";
  import meetups from "./Meetups/meetups-store.js";

  //let loadedMeetups = [];

  let editMode = null;
  let editedId = null;
  let page = "overview";
  let pageData = {};

  fetch("http://localhost:9000/api/meetups")
    .then(res => {
      if(!res.ok) {
        throw new Error('Issue fetching meetups')
      }
      return res.json();
    })
    .then(data => {
      const loadedMeetups = [];
      for (const key in data.data) {
        loadedMeetups.push({
          ...data.data[key],
         id:key
       })
      }
      //console.log(loadedMeetups)

      meetups.setMeetups(loadedMeetups)
    })
    .catch(err => console.log(err));

  const savedMeetup = () => {
    editMode = null;
    editedId = null;
  };
  const cancelEdit = () => {
    editMode = null;
    editedId = null;
  };
  const showDetails = e => {
    page = "details";
    pageData.id = e.detail;
  };

  const close = () => {
    page = "overview";
    pageData.id = {};
  };

  const startEdit = e => {
    editMode = "edit";
    editedId = e.detail;
  };
</script>

<style>
  main {
    margin-top: 5rem;
  }
</style>

<Header />

<main>
  {#if page === 'overview'}
    {#if editMode === 'edit'}
      <EditMeetup
        id={editedId}
        on:addmeetup={savedMeetup}
        on:cancel={cancelEdit} />
    {/if}
    <MeetUpGrid
      meetups={$meetups}
      on:showdetails={showDetails}
      on:edit={startEdit}
      on:add={() => (editMode = 'edit')} />
  {:else}
    <MeetupDetail id={pageData.id} on:close={close} />
  {/if}
</main>
