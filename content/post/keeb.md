---
title: "A super tiny, practical and cheap keyboard build"
date: 2021-08-15T14:54:56+01:00
tags: ["programming", "keyboard", "linux", "dotfiles"]
categories: ["keyboards"]
---
For a long while now I've been interested in mechanical keyboards. But for the most part I hadn't had the money to take part, boards tended to be expensive and as an unemployed student I needed to find something on the cheap. My first stop was a cheap ten keyless board from Amazon. It had some form of nockoff cherry blues. It was a good daily driver for the rather lengthy time I was using it but i decided it was time for change. As I was looking around through many sources for mechanical keyboards (/r/mk and /g/mgk) I came across the olkb planck. The planck is a special looking keyboard, its described as a 40% ortholinear board. 

40% means that it is only forty percent of the size of a standard keyboard. 

Here are some of the most common layouts just so you can see the comparative size of the planck:

- Ten keyless

![/images/tkl.png](/images/tkl.png)

- 65%

![/images/65.png](/images/65.png)

- 40% - the size of the planck

![images/40.png](/images/40.png)

- 30%

![images/30.png](/images/30.png)

The next of the two criteria that needs describing is what ortholinear means. So forgive me while we do some etymology. 'ortho' is a prefix meaning straight, right or proper derived from the Greek word orthós. Linear is a Latin derived word which in this case relates to a relationship to lines. Looking at the difference between a staggered layout this will probably become more apparent

This is a staggered 40% layout:

![images/40%201.png](/images/40%201.png)

Whereas this is an ortho 40% layout:

![images/planck.png](/images/planck.png)

As you can see an ortho layout sets all the keys in a grid kind of like an ice cube tray

I knew I had to have one but then I looked at the price, USD245 ?!?! For a small 'affordable' mechanical keyboard that felt extortionate. So I took it upon myself to build a cheap 40% ortholinear keyboard. After having a look around the internet I found ai03's contra. It was just what i was looking for:

- Cheap
- Open source
- Not restricted to a group buy

So I grabbed a copy of the gerber files and sent them off to JLCPCB, a  PCB fabrication company in china. Keeping in the theme of cheap the total for 5 PCBs was 12GBP including shipping. 

In the time they took to fabricate I had to choose what switches I wanted. My first manufacturer ruled out was Cherry, they make good switches but are just a little outside of my price range. Next up was Gateron who tend to be the top choice for cheap keyswitches but there was little to no stock wherever I looked which I can imagine is in most part due to COVID-19, so my last option was Kailh. Kailh has a rather wide range of switches so picking wasnt easy. I knew I wanted linear switches so that narrowed my options down a bit. After a bit of looking around I came across the Berry Heavy Switches, the heavy part alluding to the 70g actuation and 95g bottom out forces making them pretty heavy in comparison to something like cherry blue switches which actuate at 60g. And I will tell you know you can *feel* it. 

Once all was ordered (PCB, switches, diodes and microcontroller) it was time to wait, and wait I did. Which gave me ample time to figure out a solution for a plate. As the berrys are plate mount switches it would be in advisable at best to go without a plate. My original idea was to get the plate laser cut out of some 3mm acrylic. But, laser cutting is expensive. That being said I did see if I could get it done at school but my attempt was met with a resounding no response. That left me with one final option which was asking a friend if he could 3D print it for me which he very generously obliged. Leaving me with just the base to figure out. As I knew there was some wood lying around the house I decided id put something together out of that.

Once all the pieces arrived it was time to get building. The guide on the website of the original contra group buy was very helpful and I managed to get the whole thing together rather fast. Which meant it was time to come up with a keymap and flash it. The contra is QMK compatible meaning you can completely define your own keymap which is incredibly useful for defining extra functionally. This is especially usefull for someone like me who's daily computing tasks really heavily. An example of this is switching workspaces within dwm (my choice of tiling window manager), the key combination for this is the mod key plus the number key associated with that workspace, so switching to workspace one would be associated with the key bind `mod+1`, but this isn't a possible combination with one layer due to not having a number row. My solution for this is to use a special QMK key which combines two key presses into one key and placing that key onto a separate layer. There are a lot more cases like this so rather than explain all of them you can take a look at my keymap below:

![images/keymap.png](/images/keymap.png)

If you have any specific questions about some of the keymap choices they may be rationalized by looking at my [dotfiles](https://github.com/manfromth3m0oN/dots)
or my dwm build [`config.h`](https://github.com/manfromth3m0oN/dwm-bsd/blob/master/config.h)

So now the important bit: is it any good?

In short: yes

In longer: Yes, a smaller keyboard has got a bunch of advantages. One namely benefit is the more space on ones desk. I have a pretty large desk so that's not a huge issue for me. The next big one is the reduced stress on your hands. I have found a lot of truth to this, with the reduced stretching involved
