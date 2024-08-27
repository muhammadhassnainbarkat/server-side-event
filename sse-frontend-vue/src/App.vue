<script setup>
import {onUnmounted, ref} from "vue";

const error = ref(null)
const messages = ref([]);
const eventSource = ref(null);
const clientId = "1";

const subscribe = () => {

  if(eventSource.value) {return}

  eventSource.value = new EventSource("http://localhost:8080/event/subscribe?clientId=" + clientId);
  eventSource.value.onmessage = (event) => {
    messages.value.push(event.data)
  }

  eventSource.value.onerror = (event) => {
    this.error.value = event.error
  }

}

const disconnect = () => {
  if(eventSource.value){
    fetch('http://localhost:8080/event/unsubscribe?clientId=' + clientId)
        .then(() => {
          eventSource.value.close();
          eventSource.value = null;
          messages.value = [];
        })
  }
}

onUnmounted(()=> {
  disconnect();
})

</script>

<template>
  <main>
    <div class="container">
      <h1> Server side event</h1>

      <p v-if="error" class="error">{{error}}</p>
      <div class="button-group">
        <button @click="subscribe">Subscribe</button>
        <button @click="disconnect">Unsubscribe</button>
      </div>
      <div v-for="(message, index) in messages" :key="index" class="message-list">
        <p class="message">{{ message }}</p>
      </div>

      </div>
  </main>
</template>

<style scoped>
.container {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  text-align: center;
  background-color: #f9f9f9;
  border-radius: 10px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
}

h1 {
  font-size: 2rem;
  color: #333;
  margin-bottom: 20px;
}

.error {
  color: #ff4d4f;
  font-size: 1.2rem;
  margin-bottom: 15px;
}

.message-list {
  max-height: 300px;
  overflow-y: auto;
  margin-bottom: 5px;
}

.message {
  background-color: #e0f7fa;
  padding: 10px;
  margin: 5px 0;
  border-radius: 5px;
  font-size: 1.1rem;
  color: #00796b;
}

.button-group {
  display: flex;
  justify-content: center;
  gap: 20px;
}

button {
  padding: 10px 20px;
  font-size: 1rem;
  color: white;
  background-color: #007bff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #0056b3;
}

button:active {
  background-color: #004085;
}
</style>
