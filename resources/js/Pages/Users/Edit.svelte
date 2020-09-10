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
        Id: data.Id,
        FirstName: data.FirstName || '',
        LastName: data.LastName || '',
        Email: data.Email || '',
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

<Helmet title={`${values.FirstName} ${values.LastName}`} />

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
                {values.FirstName} {values.LastName}
            </h1>

            {#if data.PhotoPath}
                <img class="block w-8 h-8 rounded-full ml-4" src={data.PhotoPath} alt={data.FirstName + " " + data.LastName} />
            {/if}
        </div>

        {#if data.DeletedAt}
            <TrashedMessage>This user has been deleted.</TrashedMessage>
        {/if}

        {#if errors && errors["general"] }
            <div class="form-error">{errors["general"]}</div>
        {/if}

        <div class="bg-white rounded shadow overflow-hidden max-w-3xl">
            <form on:submit|preventDefault={handleSubmit}>
                <div class="p-8 -mr-6 -mb-8 flex flex-wrap">
                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="First Name"
                        name="FirstName"
                        errors={errors["User.FirstName"]}
                        value={values.FirstName}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Last Name"
                        name="LastName"
                        errors={errors["User.LastName"]}
                        value={values.LastName}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Email"
                        name="Email"
                        type="email"
                        errors={errors["User.Email"]}
                        value={values.Email}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Password"
                        name="password"
                        type="password"
                        errors={errors["User.Password"]}
                        value={values.password}
                        onChange={handleChange}
                    />

                    <SelectInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Owner"
                        name="Owner"
                        errors={errors["User.Owner"]}
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
