<script lang="ts">
    export let title: string;
    export let amount: number;
    export let frequency: string;
    export let category: string ;
    export let description: string;
</script>

<tr>
    <td>{title}</td>
    <td>{amount}</td>
    <td>{frequency}</td>
    <td>{category}</td>
    <td>{description}</td>
    <td>
        <button class="btn btn-primary">Edit</button>
        <button class="btn btn-danger">Delete</button>
    </td>
</tr>