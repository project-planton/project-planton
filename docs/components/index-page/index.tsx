import { Feature, Features } from '@components/features'
import Image from 'next/image'
import { Link } from 'nextra-theme-docs'
import manifestsCardDark from 'public/images/landing/manifests-dark.png'
import pulumiCodeDark from 'public/images/landing/pulumi-code-block.png'
import manifestsCard from 'public/images/landing/manifests-light.png'
import styles from './index.module.css'


const YouTubeEmbed = ({ videoId }) => {
    return (
        <div style={{ position: 'relative', paddingBottom: '56.25%', height: 0, overflow: 'hidden', maxWidth: '100%', background: '#000', borderRadius: '1.5rem' }}>
            <iframe
                src={`https://www.youtube.com/embed/${videoId}`}
                style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%' }}
                frameBorder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                allowFullScreen
                title="Embedded YouTube"
            />
        </div>
    );
};

export const IndexPage = () => (
    <div className="home-content">
        <div className="content-container grid-box">
            <div>
                <h1 className="headline">
                    OpenSource Multi-Cloud <br className="sm:block hidden"/>
                    Deployment Framework
                </h1>
                <p className="subtitle">
                    Built w/ everything you love from{' '}
                    <Link href="https://github.com/kubernetes/design-proposals-archive/blob/main/architecture/resource-management.md" className="text-current">
                        Kubernetes Resource Model (KRM)
                    </Link>{' , '}
                    <br className="max-md:hidden"/>
                    <Link href="https://protobuf.dev/" className="text-current">
                        Protobuf
                    </Link>{' , '}
                    <Link href="https://github.com/bufbuild/buf" className="text-current">
                        Buf
                    </Link>{' & '}
                    <Link href="https://github.com/pulumi/pulumi" className="text-current">
                        Pulumi
                    </Link>
                </p>
                <p className="subtitle">
                    <Link className={styles.cta} href="/docs">
                        Get started <span>→</span>
                    </Link>
                </p>
            </div>
            <div className="video-container">
                <YouTubeEmbed videoId="rbdMs56uGNI?si=MlGaV1xZukYlYmpX" />
            </div>
        </div>
        <style jsx>{`
            .grid-box {
                display: flex;
                justify-content: space-between;
            }
            .video-container {
                width: 40%;
                margin: auto 0;
                margin-top: 4rem;
            }
            .content-container {
                max-width: 90rem;
                padding-left: max(env(safe-area-inset-left), 1.5rem);
                padding-right: max(env(safe-area-inset-right), 1.5rem);
                margin: 0 auto;
            }

            .features-container {
                margin: 8rem 0 0;
                padding: 4rem 0;
                background-color: #f3f4f6;
                border-bottom: 1px solid #e5e7eb;
            }

            .features-container .content-container {
                margin-top: -8rem;
            }

            :global(.dark) .features-container {
                background-color: #000;
                border-bottom: 1px solid rgb(38, 38, 38);
            }

            .headline {
                display: inline-flex;
                font-size: 3.125rem;
                font-size: min(4.375rem, max(8vw, 2.5rem));
                font-weight: 700;
                font-feature-settings: initial;
                letter-spacing: -0.12rem;
                margin-left: -0.2rem;
                margin-top: 3.4rem;
                line-height: 1.1;
                background-image: linear-gradient(146deg, #000, #757a7d);
                -webkit-background-clip: text;
                -webkit-text-fill-color: transparent;
                background-clip: text;
            }

            :global(.dark) .headline {
                background-image: linear-gradient(146deg, #fff, #757a7d);
            }

            .subtitle {
                font-size: 1.3rem;
                font-size: min(1.3rem, max(3.5vw, 1.2rem));
                font-feature-settings: initial;
                line-height: 1.6;
            }

            :global(#docs-card) {
                color: #fff;
                text-shadow: 0 0 1rem rgba(0, 0, 0, 0.1);
            }

            :global(#docs-card img) {
                object-fit: cover;
                object-position: left;
                position: absolute;
                left: 0;
                right: 0;
                top: 0;
                bottom: 0;
                height: 100%;
                z-index: 0;
                user-select: none;
                pointer-events: none;
            }

            :global(#docs-card img:nth-child(2)) {
                display: none;
            }

            :global(.dark #docs-card img:nth-child(2)) {
                display: initial;
            }

            :global(.dark #docs-card img:nth-child(1)) {
                display: none;
            }

            :global(#highlighting-card) {
                min-height: 300px;
                background-image: linear-gradient(to top, transparent, #fff 50%),
                url(/assets/syntax-highlighting.svg);
                background-size: 634px;
                background-position: -6px calc(100% + 4px);
                background-repeat: no-repeat;
            }

            :global(.dark #highlighting-card) {
                background-image: linear-gradient(to top, transparent, #202020 50%),
                url(/assets/syntax-highlighting.svg);
            }

            :global(.feat-darkmode) {
                min-height: 300px;
            }

            :global(.feat-darkmode h3) {
                font-size: 48px;
            }

            :global(#search-card) {
                display: flex;
                flex-direction: column;
                justify-content: center;
            }

            :global(#search-card p) {
                max-width: 320px;
            }

            :global(#search-card video) {
                position: absolute;
                right: 0;
                top: 24px;
                height: 430px;
                pointer-events: none;
                max-width: 60%;
            }

            :global(#fs-card) {
                min-height: 240px;
            }

            :global(#fs-card h3) {
                text-align: left;
                width: min(300px, 41%);
                min-width: 155px;
            }

            :global(#a11y-card) {
                background-image: url(/assets/high-contrast.png);
                background-position: -160px 160px;
            }

            @media screen and (max-width: 1300px) {
                :global(#a11y-card) {
                    background-image: linear-gradient(to bottom, white, transparent),
                    url(/assets/high-contrast.png);
                }
            }

            @media screen and (max-width: 1200px) {
                :global(#highlighting-card) {
                    aspect-ratio: auto;
                }

                :global(.feat-darkmode h3) {
                    font-size: 4vw;
                    font-size: min(48px, max(4vw, 30px));
                }

                :global(#search-card video) {
                    aspect-ratio: 787/623;
                    height: auto;
                }

                .headline {
                    letter-spacing: -0.08rem;
                }
            }

            @media screen and (max-width: 1024px) {
                :global(#docs-card) {
                    aspect-ratio: 135/86;
                }

                :global(#search-card) {
                    aspect-ratio: 8/3;
                }

                :global(#search-card h3) {
                    text-align: left;
                }

                :global(#highlighting-card) {
                    background-size: 136%;
                }

                :global(#a11y-card) {
                    background-image: url(/assets/high-contrast.png);
                    background-position: center 160px;
                }
            }

            @media screen and (max-width: 768px) {
                :global(#docs-card) {
                    min-height: 348px;
                    width: 100%;
                    aspect-ratio: auto;
                }

                :global(#docs-card img) {
                    object-position: -26px 0;
                    width: 250%;
                    max-width: initial;
                }
            }

            @media screen and (max-width: 640px) {
                :global(#search-card) {
                    aspect-ratio: 2.5/2;
                    justify-content: flex-start;
                    align-items: stretch;
                    min-height: 350px;
                }

                :global(#search-card h3) {
                    text-align: center;
                }

                :global(#search-card p) {
                    max-width: 100%;
                }

                :global(#search-card video) {
                    position: relative;
                    margin: 0.75em -1.75em 0;
                    max-width: calc(100% + 3.5em);
                }

                :global(.dark #search-card video) {
                    mix-blend-mode: lighten;
                }
            }
        `}</style>
        <div className="features-container">
            <div className="content-container">
                <Features>
                    <Feature
                        index={0}
                        large
                        centered
                        id="docs-card"
                        href="/docs/api/design"
                    >
                        <Image src={manifestsCard} alt="Background" loading="eager"/>
                        <Image src={manifestsCardDark} alt="Background (Dark)" loading="eager"/>
                        <h3>
                            Everything is a Simple Manifest
                        </h3>
                    </Feature>
                    <Feature index={1} centered href="/docs/core-tenets-02-pulumi-modules">
                        <h3>
                            Pulumi Modules <br className="show-on-mobile"/>
                            <span className="font-light">for all Deployment Components</span>
                        </h3>
                        <p className="text-left mb-8">
                            Every API Resource has a well-written and customizable{' '}
                            <Link href="https://github.com/orgs/plantoncloud/repositories?q=pulumi-module">
                                Pulumi Module
                            </Link>{' '}
                            configured as default.
                        </p>
                        <Image src={pulumiCodeDark} alt="Background (Dark)" loading="eager"/>
                    </Feature>
                    <Feature
                        index={7}
                        large
                        id="fs-card"
                        style={{
                            color: 'white',
                            backgroundImage:
                                'url(/images/landing/pulumi-up.png), url(/images/landing/cli-tie-bg.png)',
                            backgroundSize: '140%, 180%',
                            backgroundPosition: '130px -8px, top',
                            backgroundRepeat: 'no-repeat',
                            textShadow: '0 1px 6px rgb(38 59 82 / 18%)',
                            aspectRatio: '1.765'
                        }}
                        href="/docs/docs-theme/page-configuration"
                    >
                        <h3>
                            A CLI ties it all, <br/>
                            to deliver the Magic!!
                        </h3>
                    </Feature>
                    <Feature
                        index={8}
                        style={{
                            backgroundSize: 750,
                            backgroundRepeat: 'no-repeat',
                            minHeight: 288
                        }}
                    >
                        <h3>Opinionated only at API</h3>
                        <p>
                            ProjectPlanton avoids unneeded abstractions <br className="show-on-mobile"/>
                             for <b>Maximum Customizability</b> and <b>Flexibility</b>.
                        </p>
                    </Feature>
                    {/*<Feature*/}
                    {/*    index={2}*/}
                    {/*    id="highlighting-card"*/}
                    {/*    href="/docs/core-tenets-03-cli"*/}
                    {/*>*/}
                    {/*    <h3>*/}
                    {/*        Advanced syntax <br className="show-on-mobile"/>*/}
                    {/*        highlighting solution*/}
                    {/*    </h3>*/}
                    {/*    <p>*/}
                    {/*        Performant and reliable build-time syntax highlighting powered by{' '}*/}
                    {/*        <Link href="https://shiki.style">Shiki</Link>.*/}
                    {/*    </p>*/}
                    {/*</Feature>*/}
                    {/*<Feature index={3} href="/docs/guide/i18n">*/}
                    {/*    <h3>*/}
                    {/*        I18n as easy as <br className="show-on-mobile"/>*/}
                    {/*        creating new files*/}
                    {/*    </h3>*/}
                    {/*    <p className="mb-4">*/}
                    {/*        Name your page files with locales suffixed, Nextra and Next.js*/}
                    {/*        will do the rest for you.*/}
                    {/*    </p>*/}
                    {/*</Feature>*/}
                    {/*<Feature*/}
                    {/*    index={4}*/}
                    {/*    centered*/}
                    {/*    style={{*/}
                    {/*        display: 'flex',*/}
                    {/*        flexDirection: 'column',*/}
                    {/*        alignItems: 'center',*/}
                    {/*        justifyContent: 'center',*/}
                    {/*        backgroundImage: 'url(/assets/gradient-bg.jpeg)',*/}
                    {/*        backgroundSize: 'cover',*/}
                    {/*        backgroundPosition: 'center',*/}
                    {/*        color: '#fff'*/}
                    {/*    }}*/}
                    {/*    href="/docs/guide/markdown"*/}
                    {/*>*/}
                    {/*    <svg*/}
                    {/*        width="70%"*/}
                    {/*        viewBox="0 0 69 29"*/}
                    {/*        fill="none"*/}
                    {/*        xmlns="http://www.w3.org/2000/svg"*/}
                    {/*        style={{*/}
                    {/*            filter: 'drop-shadow(0 2px 10px rgba(0, 0, 0, .1))'*/}
                    {/*        }}*/}
                    {/*    >*/}
                    {/*        <path*/}
                    {/*            fillRule="evenodd"*/}
                    {/*            clipRule="evenodd"*/}
                    {/*            d="M66.375 0.375H2.625C1.38236 0.375 0.375 1.38236 0.375 2.625V25.875C0.375 27.1176 1.38236 28.125 2.625 28.125H66.375C67.6176 28.125 68.625 27.1176 68.625 25.875V2.625C68.625 1.38236 67.6176 0.375 66.375 0.375ZM23.625 5.75368V9.375V21.875H20.625V12.9963L16.1857 17.4357L15.125 18.4963L14.0643 17.4357L9.75 13.1213V22H6.75V9.5V5.87868L9.31066 8.43934L15.125 14.2537L21.0643 8.31434L23.625 5.75368ZM29.5607 12.5643L33.75 16.7537V5.375H36.75V16.7537L40.9393 12.5643L43.0607 14.6857L36.3107 21.4357L35.25 22.4963L34.1893 21.4357L27.4393 14.6857L29.5607 12.5643ZM62.3105 19.5592L56.1228 13.3736L62.4357 7.06066L60.3143 4.93934L54.0011 11.2526L47.6855 4.93916L45.5645 7.06084L51.8798 13.3739L45.6893 19.5643L47.8107 21.6857L54.0014 15.4949L60.1895 21.6808L62.3105 19.5592Z"*/}
                    {/*            fill="#fff"*/}
                    {/*        />*/}
                    {/*    </svg>*/}
                    {/*    <p*/}
                    {/*        style={{*/}
                    {/*            textShadow: '0 2px 4px rgb(0 0 0 / 20%)'*/}
                    {/*        }}*/}
                    {/*    >*/}
                    {/*        <Link href="https://mdxjs.com/blog/v3" className="text-current">*/}
                    {/*            MDX 3*/}
                    {/*        </Link>{' '}*/}
                    {/*        lets you use Components inside Markdown,{' '}*/}
                    {/*        <br className="hide-medium"/>*/}
                    {/*        with huge performance boost since v1.*/}
                    {/*    </p>*/}
                    {/*</Feature>*/}
                    {/*<Feature*/}
                    {/*    index={5}*/}
                    {/*    centered*/}
                    {/*    className="feat-darkmode"*/}
                    {/*    style={{*/}
                    {/*        display: 'flex',*/}
                    {/*        alignItems: 'center',*/}
                    {/*        justifyContent: 'center'*/}
                    {/*    }}*/}
                    {/*>*/}
                    {/*    <motion.div*/}
                    {/*        animate={{*/}
                    {/*            backgroundPosition: [*/}
                    {/*                '0% 0%',*/}
                    {/*                '50% 40%',*/}
                    {/*                '50% 40%',*/}
                    {/*                '100% 100%'*/}
                    {/*            ],*/}
                    {/*            backgroundImage: [*/}
                    {/*                'radial-gradient(farthest-corner, #e2e5ea, #e2e5ea)',*/}
                    {/*                'radial-gradient(farthest-corner, #06080a, #e2e5ea)',*/}
                    {/*                'radial-gradient(farthest-corner, #06080a, #e2e5ea)',*/}
                    {/*                'radial-gradient(farthest-corner, #e2e5ea, #e2e5ea)'*/}
                    {/*            ]*/}
                    {/*        }}*/}
                    {/*        transition={{*/}
                    {/*            backgroundPosition: {*/}
                    {/*                times: [0, 0.5, 0.5, 1],*/}
                    {/*                repeat: Infinity,*/}
                    {/*                duration: 10,*/}
                    {/*                delay: 1*/}
                    {/*            },*/}
                    {/*            backgroundImage: {*/}
                    {/*                times: [0, 0.2, 0.8, 1],*/}
                    {/*                repeat: Infinity,*/}
                    {/*                duration: 10,*/}
                    {/*                delay: 1*/}
                    {/*            }*/}
                    {/*        }}*/}
                    {/*        style={{*/}
                    {/*            position: 'absolute',*/}
                    {/*            top: 0,*/}
                    {/*            left: 0,*/}
                    {/*            width: '100%',*/}
                    {/*            height: '100%',*/}
                    {/*            backgroundImage:*/}
                    {/*                'radial-gradient(farthest-corner, #06080a, #e2e5ea)',*/}
                    {/*            backgroundSize: '400% 400%',*/}
                    {/*            backgroundRepeat: 'no-repeat'*/}
                    {/*        }}*/}
                    {/*    />*/}
                    {/*    <motion.h3*/}
                    {/*        animate={{*/}
                    {/*            color: ['#dae5ff', '#fff', '#fff', '#dae5ff']*/}
                    {/*        }}*/}
                    {/*        transition={{*/}
                    {/*            color: {*/}
                    {/*                times: [0.25, 0.35, 0.7, 0.8],*/}
                    {/*                repeat: Infinity,*/}
                    {/*                duration: 10,*/}
                    {/*                delay: 1*/}
                    {/*            }*/}
                    {/*        }}*/}
                    {/*        style={{*/}
                    {/*            position: 'relative',*/}
                    {/*            mixBlendMode: 'difference'*/}
                    {/*        }}*/}
                    {/*    >*/}
                    {/*        Dark <br/>*/}
                    {/*        mode <br/>*/}
                    {/*        included*/}
                    {/*    </motion.h3>*/}
                    {/*</Feature>*/}
                    {/*<Feature*/}
                    {/*    index={6}*/}
                    {/*    large*/}
                    {/*    id="search-card"*/}
                    {/*    href="/docs/docs-theme/theme-configuration#search"*/}
                    {/*>*/}
                    {/*    <div style={{zIndex: 2}}>*/}
                    {/*        <h3>*/}
                    {/*            Full-text search, <br/>*/}
                    {/*            zero-config needed*/}
                    {/*        </h3>*/}
                    {/*        <p>*/}
                    {/*            Nextra indexes your content automatically at build-time and*/}
                    {/*            performs incredibly fast full-text search via{' '}*/}
                    {/*            <Link href="https://github.com/nextapps-de/flexsearch">*/}
                    {/*                FlexSearch*/}
                    {/*            </Link>*/}
                    {/*            .*/}
                    {/*        </p>*/}
                    {/*    </div>*/}
                    {/*    <div*/}
                    {/*        className="absolute size-full inset-0 hidden sm:block bg-[linear-gradient(to_right,white_250px,_transparent)] dark:bg-[linear-gradient(to_right,#202020_250px,_transparent)] z-[1]"/>*/}
                    {/*    <video*/}
                    {/*        autoPlay*/}
                    {/*        loop*/}
                    {/*        muted*/}
                    {/*        playsInline*/}
                    {/*        className="nextra-focus dark:hidden block"*/}
                    {/*    >*/}
                    {/*        <source src="/assets/search.mp4" type="video/mp4"/>*/}
                    {/*    </video>*/}
                    {/*    <video*/}
                    {/*        autoPlay*/}
                    {/*        loop*/}
                    {/*        muted*/}
                    {/*        playsInline*/}
                    {/*        className="nextra-focus dark:block hidden -translate-x-4"*/}
                    {/*    >*/}
                    {/*        <source src="/assets/search-dark.mp4" type="video/mp4"/>*/}
                    {/*    </video>*/}
                    {/*</Feature>*/}
                    {/*<Feature index={9} href="/docs/guide/ssg">*/}
                    {/*    <h3>*/}
                    {/*        Hybrid rendering, <br/>*/}
                    {/*        next generation*/}
                    {/*    </h3>*/}
                    {/*    <p className="mr-6">*/}
                    {/*        You can leverage the hybrid rendering power from Next.js with your*/}
                    {/*        Markdown content including{' '}*/}
                    {/*        <Link href="https://nextjs.org/docs/basic-features/pages#static-generation-recommended">*/}
                    {/*            SSG*/}
                    {/*        </Link>*/}
                    {/*        ,{' '}*/}
                    {/*        <Link href="https://nextjs.org/docs/basic-features/pages#server-side-rendering">*/}
                    {/*            SSR*/}
                    {/*        </Link>*/}
                    {/*        , and{' '}*/}
                    {/*        <Link*/}
                    {/*            href="https://nextjs.org/docs/basic-features/data-fetching/incremental-static-regeneration">*/}
                    {/*            ISR*/}
                    {/*        </Link>*/}
                    {/*        .*/}
                    {/*    </p>*/}
                    {/*</Feature>*/}
                    {/*<Feature index={10} large>*/}
                    {/*    <h3>And more...</h3>*/}
                    {/*    <p>*/}
                    {/*        SEO / RTL Layout / Pluggable Themes / Built-in Components / Last*/}
                    {/*        Git Edit Time / Multi-Docs...*/}
                    {/*        <br/>A lot of new possibilities to be explored.*/}
                    {/*    </p>*/}
                    {/*    <p className="subtitle">*/}
                    {/*        <Link className="no-underline" href="/docs">*/}
                    {/*            Start using Nextra →*/}
                    {/*        </Link>*/}
                    {/*    </p>*/}
                    {/*</Feature>*/}
                </Features>
            </div>
        </div>
    </div>
)
