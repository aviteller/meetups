import { writable } from "svelte/store";

const meetups = writable([
]);

const customMeetupsStore = {
  subscribe: meetups.subscribe,
  setMeetups: (meetupArray) => {

    meetups.set(meetupArray);
  },
  addMeetup: meetUpData => {
    const newMeetup = {
      ...meetUpData
    };
    meetups.update(items => [newMeetup, ...items]);
  },
  toggleLike: id => {
    meetups.update(items => {
      const updatedMeetup = { ...items.find(m => m.ID === id) };
      updatedMeetup.isLiked = !updatedMeetup.isLiked;
      const meetupIndex = items.findIndex(m => m.ID === id);
      const updatedMeetups = [...items];
      updatedMeetups[meetupIndex] = updatedMeetup;
      return updatedMeetups;
    });
  },
  updateMeetup: (id, meetUpData) => {
    meetups.update(items => {
      const meetupIndex = items.findIndex(m => m.id === id);
      const updatedMeetup = { ...items[meetupIndex], ...meetUpData };
      const updatedMeetups = [...items];
      updatedMeetups[meetupIndex] = updatedMeetup;
      return updatedMeetups;
    });
  },
  removeMeetup: id => {
    meetups.update(items => {
      return items.filter(i => i.ID !== id);
    });
  }
};

export default customMeetupsStore;
