<script lang="ts">
  export let message: string;
</script>

<error>
  <div class="success-box">
    <span class="success-message">✅ {message}</span>
  </div>
</error>

<style>
  .success-box {
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
    background-color: #9affcc;
    border: 1px solid #048000;
    border-radius: 5px;
    color: #0bc215;
    font-size: 16px;
    padding: 10px;
    margin: 10px;
  }

  @keyframes slide-in {
    100% {
      left: 20px;
    }
  }

  .success-message {
    font-weight: bold;
  }
</style>
