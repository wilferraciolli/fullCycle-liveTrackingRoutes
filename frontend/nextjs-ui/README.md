## Next JS fronent
This project will be used to be the front end of the routing track, it uses NextJS framework, with React and tailwind.
It will be mainly SSR.
this also will need to integrate with Google APIs to display the map to the user.
![1-overview.png](../images/1-overview.png)

# Next Js and SSR
The idea is to have server side rended pages rather than a traditional SPA. the server will render the page and send the page to the user
only additional JS will be added in case the page needs to.
![2-next-js-architecture.png](../images/2-next-js-architecture.png)


# Create the project 
run `npx create-next-app@latest nestjs-ui` to create this project

# Running the app
`npm run dev` - default is port 300 but it can be changeg

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
