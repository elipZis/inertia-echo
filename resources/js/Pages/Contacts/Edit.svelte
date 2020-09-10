<script>
    import { Inertia } from '@inertiajs/inertia';
    import { InertiaLink, page } from '@inertiajs/inertia-svelte';
    import Helmet from '@/Shared/Helmet.svelte';
    import Layout from '@/Shared/Layout.svelte';
    import DeleteButton from '@/Shared/DeleteButton.svelte';
    import LoadingButton from '@/Shared/LoadingButton.svelte';
    import TextInput from '@/Shared/TextInput.svelte';
    import SelectInput from '@/Shared/SelectInput.svelte';
    import TrashedMessage from '@/Shared/TrashedMessage.svelte';
    import { toFormData } from '@/utils';

    const route = window.route;

    let { data, organizations } = $page;
    $: data = $page.data;
    $: organizations = $page.organizations ?? [];
    $: errors = $page.errors ?? [];

    let sending = false;
    let values = {
        Id: data.Id,
        FirstName: data.FirstName || '',
        LastName: data.LastName || '',
        OrganizationId: data.OrganizationId || '',
        Email: data.Email || '',
        Phone: data.Phone || '',
        Address: data.Address || '',
        City: data.City || '',
        Region: data.Region || '',
        Country: data.Country || '',
        PostalCode: data.PostalCode || ''
    };

    function handleChange({ target: { name, value } }) {
        values ={
            ...values,
            [name]: value
        };
    }

    function handleSubmit(e) {
        sending = true;
        console.log(values);
        const formData = toFormData(values);
        Inertia.post(route('contacts.update', {contact: data.Id}), formData).then(() => sending = false);
    }

    function destroy() {
        if (confirm('Are you sure you want to delete this contact?')) {
            Inertia.delete(route('contacts.destroy', {contact: data.Id}));
        }
    }
</script>

<Helmet title={`${values.FirstName} ${values.LastName}`} />

<Layout>
    <div>
        <h1 class="mb-8 font-bold text-3xl">
            <InertiaLink
                href={route('contacts')}
                class="text-indigo-600 hover:text-indigo-700"
            >
                Contacts
            </InertiaLink>

            <span class="text-indigo-600 font-medium mx-2">/</span>
            {values.FirstName} {values.LastName}
        </h1>

        {#if data.DeletedAt}
            <TrashedMessage>This contact has been deleted.</TrashedMessage>
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
                        errors={errors["Contact.FirstName"]}
                        value={values.FirstName}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Last Name"
                        name="LastName"
                        errors={errors["Contact.LastName"]}
                        value={values.LastName}
                        onChange={handleChange}
                    />

                    <SelectInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Organization"
                        name="OrganizationId"
                        errors={errors["Contact.OrganizationId"]}
                        value={values.OrganizationId}
                        onChange={handleChange}
                    >
                        <option value=""></option>
                        {#each organizations as { Id, Name } (Id)}
                            <option value={`${Id}`}>{Name}</option>
                        {/each}
                    </SelectInput>

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Email"
                        name="Email"
                        type="email"
                        errors={errors["Contact.Email"]}
                        value={values.Email}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Phone"
                        name="Phone"
                        type="text"
                        errors={errors["Contact.Phone"]}
                        value={values.Phone}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Address"
                        name="Address"
                        type="text"
                        errors={errors["Contact.Address"]}
                        value={values.Address}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="City"
                        name="City"
                        type="text"
                        errors={errors["Contact.City"]}
                        value={values.City}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Province/State"
                        name="Region"
                        type="text"
                        errors={errors["Contact.Region"]}
                        value={values.Region}
                        onChange={handleChange}
                    />

                    <SelectInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Country"
                        name="Country"
                        errors={errors["Contact.Country"]}
                        value={values.Country}
                        onChange={handleChange}
                    >
                        <option value=""></option>
                        <option value="CA">Canada</option>
                        <option value="US">United States</option>
                    </SelectInput>

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Postal Code"
                        name="PostalCode"
                        type="text"
                        errors={errors["Contact.PostalCode"]}
                        value={values.PostalCode}
                        onChange={handleChange}
                    />
                </div>

                <div class="px-8 py-4 bg-gray-100 border-t border-gray-200 flex items-center">
                    {#if !data.DeletedAt}
                        <DeleteButton onDelete={destroy}>Delete Contact</DeleteButton>
                    {/if}

                    <LoadingButton
                        loading={sending}
                        type="submit"
                        className="btn-indigo ml-auto"
                    >
                        Update Contact
                    </LoadingButton>
                </div>
            </form>
        </div>
    </div>
</Layout>
