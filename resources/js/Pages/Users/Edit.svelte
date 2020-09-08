<script>
    import { Inertia } from '@inertiajs/inertia';
    import { InertiaLink, page } from '@inertiajs/inertia-svelte';
    import Helmet from '@/Shared/Helmet.svelte';
    import Layout from '@/Shared/Layout.svelte';
    import DeleteButton from '@/Shared/DeleteButton.svelte';
    import LoadingButton from '@/Shared/LoadingButton.svelte';
    import TextInput from '@/Shared/TextInput.svelte';
    import SelectInput from '@/Shared/SelectInput.svelte';
    import FileInput from '@/Shared/FileInput.svelte';
    import TrashedMessage from '@/Shared/TrashedMessage.svelte';
    import { toFormData } from '@/utils';

    const route = window.route;
    let { user, data } = $page;
    // $: user = data;
    $: errors = $page.errors ?? [];

    let sending = false;
    let values = {
        id: data.Id,
        first_name: data.FirstName || '',
        last_name: data.LastName || '',
        email: data.Email || '',
        password: data.Password || '',
        owner: data.Owner ? '1' : '0' || '0',
        photo: '',
    };

    function handleChange({ target: { name, value } }) {
        values = {
            ...values,
            [name]: value
        };
    }

    function handleFileChange(file) {
        values = {
            ...values,
            photo: file
        };
    }

    function handleSubmit() {
        sending = true;
        const formData = toFormData(values);
        Inertia.post(route('users.update'), formData).then(() => sending = false);
    }

    function destroy() {
        if (confirm('Are you sure you want to delete this user?')) {
            Inertia.delete(route('users.destroy', { user: data.Id }));
        }
    }
</script>

<Helmet title={`${values.first_name} ${values.last_name}`} />

<Layout>
    <div>
        <div class="mb-8 flex justify-start max-w-lg">
            <h1 class="font-bold text-3xl">
                <InertiaLink
                    href={route('users')}
                    class="text-indigo-600 hover:text-indigo-700"
                >
                    Users
                </InertiaLink>

                <span class="text-indigo-600 font-medium mx-2">/</span>
                {values.first_name} {values.last_name}
            </h1>

            {#if data.PhotoPath}
                <img class="block w-8 h-8 rounded-full ml-4" src={data.PhotoPath} alt={data.FirstName + " " + data.LastName} />
            {/if}
        </div>

        {#if data.DeletedAt}
            <TrashedMessage>This user has been deleted.</TrashedMessage>
        {/if}

        <div class="bg-white rounded shadow overflow-hidden max-w-3xl">
            <form on:submit|preventDefault={handleSubmit}>
                <div class="p-8 -mr-6 -mb-8 flex flex-wrap">
                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="First Name"
                        name="first_name"
                        errors={errors.first_name}
                        value={values.first_name}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Last Name"
                        name="last_name"
                        errors={errors.last_name}
                        value={values.last_name}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Email"
                        name="email"
                        type="email"
                        errors={errors.email}
                        value={values.email}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Password"
                        name="password"
                        type="password"
                        errors={errors.password}
                        value={values.password}
                        onChange={handleChange}
                    />

                    <SelectInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Owner"
                        name="Owner"
                        errors={errors.owner}
                        value={values.Owner}
                        onChange={handleChange}
                    >
                        <option value="1">Yes</option>
                        <option value="0">No</option>
                    </SelectInput>

<!--                    <FileInput-->
<!--                        className="pr-6 pb-8 w-full lg:w-1/2"-->
<!--                        label="Photo"-->
<!--                        name="photo"-->
<!--                        accept="image/*"-->
<!--                        errors={errors.photo}-->
<!--                        value={values.photo}-->
<!--                        onChange={handleFileChange}-->
<!--                    />-->
                </div>

                <div class="px-8 py-4 bg-gray-100 border-t border-gray-200 flex items-center">
                    {#if !data.DeletedAt}
                        <DeleteButton onDelete={destroy}>Delete User</DeleteButton>
                    {/if}

                    <LoadingButton
                        loading={sending}
                        type="submit"
                        className="btn-indigo ml-auto"
                    >
                        Update User
                    </LoadingButton>
                </div>
            </form>
        </div>
    </div>
</Layout>
