<script>
    let name = 'World';

    const GetMessage = () => {
        fetch('/api', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name })
        })
            .then(response => response.json())
            .then(data => {
              document.querySelector('.message').textContent = 'The Server says: ' + data.message;
            })
            .catch(error => {
              document.querySelector('.message').textContent = 'Error: ' + error;
            });
    };
</script>

<p class="message"></p>
<form on:submit|preventDefault={GetMessage}>
    <input type="text" name="name" placeholder="Name" bind:value={name}>
    <button type="submit">Send Message</button>
</form>
