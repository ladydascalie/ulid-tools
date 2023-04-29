<script lang="ts">
  export let message: string;
</script>

<error>
  <div class="error-box">
    <span class="error-message">🚨 {message}</span>
  </div>
</error>

<style>
  .error-box {
    position: fixed;
    bottom: 20px;
    left: -300px;
    padding: 10px;
    color: #fff;
    animation: slide-in 0.5s forwards;
    z-index: 9999;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #ffc7c7;
    border: 1px solid #ff5f5f;
    border-radius: 5px;
    color: #b40404;
    font-size: 16px;
    padding: 10px;
    margin: 10px;
  }

  @keyframes slide-in {
    100% {
      left: 20px;
    }
  }

  .error-message {
    font-weight: bold;
  }
</style>
