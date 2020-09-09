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
    import Icon from '@/Shared/Icon.svelte';

    const route = window.route;

    let { data, contacts } = $page;
    $: errors = $page.errors ?? [];
    $: data = $page.data;

    let sending = false;
    let values = {
        id: data.Id,
        name: data.Name || '',
        email: data.Email || '',
        phone: data.Phone || '',
        address: data.Address || '',
        city: data.City || '',
        region: data.Region || '',
        country: data.Country || '',
        postal_code: data.PostalCode || ''
    };

    function handleChange({ target: { name, value } }) {
        values = {
            ...values,
            [name]: value
        };
    }

    function handleSubmit() {
        sending = true;
        Inertia.post(route('organizations.update', {organization: data.Id}), values).then(() => sending = false);
    }

    function destroy() {
        if (confirm('Are you sure you want to delete this organization?')) {
            Inertia.delete(route('organizations.destroy', {organization: data.Id}));
        }
    }
</script>

<Helmet title={values.name} />

<Layout>
    <div>
        <h1 class="mb-8 font-bold text-3xl">
            <InertiaLink
                href={route('organizations')}
                class="text-indigo-600 hover:text-indigo-700"
            >
                Organizations
            </InertiaLink>

            <span class="text-indigo-600 font-medium mx-2">/</span>
            {values.name}
        </h1>

        {#if data.DeletedAt}
            <TrashedMessage>This organization has been deleted.</TrashedMessage>
        {/if}

        {#if errors && errors["general"] }
            <div class="form-error">{errors["general"]}</div>
        {/if}

        <div class="bg-white rounded shadow overflow-hidden max-w-3xl">
            <form on:submit|preventDefault={handleSubmit}>
                <div class="p-8 -mr-6 -mb-8 flex flex-wrap">
                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Name"
                        name="name"
                        errors={errors["Organization.Name"]}
                        value={values.name}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Email"
                        name="email"
                        type="email"
                        errors={errors["Organization.Email"]}
                        value={values.email}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Phone"
                        name="phone"
                        type="text"
                        errors={errors["Organization.Phone"]}
                        value={values.phone}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Address"
                        name="address"
                        type="text"
                        errors={errors["Organization.Address"]}
                        value={values.address}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="City"
                        name="city"
                        type="text"
                        errors={errors["Organization.City"]}
                        value={values.city}
                        onChange={handleChange}
                    />

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Province/State"
                        name="region"
                        type="text"
                        errors={errors["Organization.Region"]}
                        value={values.region}
                        onChange={handleChange}
                    />

                    <SelectInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Country"
                        name="country"
                        errors={errors["Organization.Country"]}
                        value={values.country}
                        onChange={handleChange}
                    >
                        <option value=""></option>
                        <option value="CA">Canada</option>
                        <option value="US">United States</option>
                    </SelectInput>

                    <TextInput
                        className="pr-6 pb-8 w-full lg:w-1/2"
                        label="Postal Code"
                        name="postal_code"
                        type="text"
                        errors={errors["Organization.PostalCode"]}
                        value={values.postal_code}
                        onChange={handleChange}
                    />
                </div>

                <div class="px-8 py-4 bg-gray-100 border-t border-gray-200 flex items-center">
                    {#if !data.DeletedAt}
                        <DeleteButton onDelete={destroy}>Delete Organization</DeleteButton>
                    {/if}

                    <LoadingButton
                        loading={sending}
                        type="submit"
                        className="btn-indigo ml-auto"
                    >
                        Update Organization
                    </LoadingButton>
                </div>
            </form>
        </div>

        <h2 class="mt-12 font-bold text-2xl">Contacts</h2>

        <div class="mt-6 bg-white rounded shadow overflow-x-auto">
            <table class="w-full whitespace-no-wrap">
                <thead>
                    <tr class="text-left font-bold">
                        <th class="px-6 pt-5 pb-4">Name</th>
                        <th class="px-6 pt-5 pb-4">City</th>
                        <th class="px-6 pt-5 pb-4" colspan="2">Phone</th>
                    </tr>
                </thead>
                <tbody>
                    {#if !contacts || contacts.length === 0}
                        <tr>
                            <td class="border-t px-6 py-4" colspan="4">
                                No contacts found.
                            </td>
                        </tr>
                    {:else}
                        {#each contacts as { Id, FirstName, LastName, Phone, City, DeletedAt } (Id)}
                            <tr class="hover:bg-gray-100 focus-within:bg-gray-100">
                                <td class="border-t">
                                    <InertiaLink
                                        href={route('contacts.edit', {contact: Id})}
                                        class="px-6 py-4 flex items-center focus:text-indigo"
                                    >
                                        {FirstName + " " + LastName}

                                        {#if DeletedAt}
                                            <Icon
                                                name="trash"
                                                className="flex-shrink-0 w-3 h-3 text-gray-400 fill-current ml-2"
                                            />
                                        {/if}
                                    </InertiaLink>
                                </td>

                                <td class="border-t">
                                    <InertiaLink
                                        tabindex="-1"
                                        href={route('contacts.edit', {contact: Id})}
                                        class="px-6 py-4 flex items-center focus:text-indigo"
                                    >
                                        {City}
                                    </InertiaLink>
                                </td>

                                <td class="border-t">
                                    <InertiaLink
                                        tabindex="-1"
                                        href={route('contacts.edit', {contact: Id})}
                                        class="px-6 py-4 flex items-center focus:text-indigo"
                                    >
                                        {Phone}
                                    </InertiaLink>
                                </td>

                                <td class="border-t w-px">
                                    <InertiaLink
                                        tabindex="-1"
                                        href={route('contacts.edit', {contact: Id})}
                                        class="px-4 flex items-center"
                                    >
                                        <Icon
                                            name="cheveron-right"
                                            className="block w-6 h-6 text-gray-400 fill-current"
                                        />
                                    </InertiaLink>
                                </td>
                            </tr>
                        {/each}
                    {/if}
                </tbody>
            </table>
        </div>
    </div>
</Layout>
