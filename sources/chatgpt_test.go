package sources

import (
	"testing"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/Axiom/util"

	gaddets "github.com/WALL-EEEEEEE/gagdets"
	core "github.com/WALL-EEEEEEE/gagdets/core"
	items "github.com/WALL-EEEEEEE/gagdets/items"
	log "github.com/sirupsen/logrus"
)

type NewsSpider struct {
	core.Spider
}

func NewNewsSpider() NewsSpider {
	spider := NewsSpider{Spider: core.NewSpider("NewsSpider", []string{})}
	spider.ITask = &spider
	return spider
}

func (spider *NewsSpider) Run() {
	var topics []items.Topic = []items.Topic{
		{
			Content: `Palestinians receive flour distributed by the United Nations on Wednesday in Khan Younis, Gaza, during a temporary truce between Hamas and Israel.Credit...Ibraheem Abu Mustafa/Reuters
International mediators were pushing on Wednesday to lock in another temporary extension of the cease-fire in Gaza between Israel and Hamas, seeing it as the best way to ease the embattled territory’s humanitarian crisis, secure the release of more Israeli captives from Gaza and slow the war’s rocketing death toll for at least a little longer.

But officials with knowledge of the talks said they also hoped that the succession of short-term pauses would pave the way toward a larger goal: negotiations over a longer-term cease-fire to bring the war to a close.

Achieving that is no easy task. Israeli officials have vowed not to stop their military campaign until Hamas’s leadership has been eliminated and the group’s military and governance infrastructure is uprooted from Gaza, objectives that remain far off. And the arrangement that has underpinned the cease-fire so far — the release of hostages from Gaza, nearly all of them women and children, in exchange for the discharge of Palestinians from Israeli prisons and the entry of more aid into Gaza — is likely to become much trickier when the parties begin negotiating to exchange captured combatants.

Hamas is believed to have captured a few dozen Israeli soldiers during the rampage it led through southern Israel on Oct. 7 that killed about 1,200 people and saw some 240 dragged into Gaza as hostages, according to Israeli officials. And Israel holds many high-profile Palestinian prisoners, including prominent members of Hamas convicted of grave crimes whose release the group has promised to pursue.

As of Tuesday, Hamas had released at least 85 hostages from Gaza, according to a New York Times tally, and Israel had released 180 Palestinians from its prisons.

With the cease-fire scheduled to expire on Thursday morning, top officials from Qatar, the United States, Egypt and Israel were meeting in Qatar in talks that were focused on extending it to allow for further exchanges. That formula has succeeded in mostly pausing the war since Friday. Each day, Israelis have watched groups of their hostages return home while Palestinians have seen their detainees released from jail, intense emotional events for each side.

The U.S. secretary of state, Antony J. Blinken, who was scheduled to travel to Israel on Thursday, said the Biden administration wants the truce to continue because it “means that more hostages will be coming home, more assistance will be getting in.”

“Clearly, that’s something we want,” Mr. Blinken told reporters in Brussels on Wednesday. “I believe it’s also something that Israel wants. They’re also intensely focused on bringing their people home.”

Speaking on condition of anonymity, an Israeli official said the country would consider further pauses as long as Hamas keeps releasing hostages.

But two people with knowledge of the talks in Qatar said that beyond yet another short-term extension, the mediators hoped to keep the war on pause for as long as possible to create conditions they hoped would allow for negotiations over a longer-term cease-fire.

One of those sources said the mediators expect that the longer the quiet lasts, the harder it will be for Israel to restart its offensive and extend it to southern Gaza, where senior Hamas leaders are believed to be hiding.

The talks have gained even more urgency because of the death toll and spiraling humanitarian crisis in Gaza. Israel’s bombardment of the territory and subsequent ground invasion have leveled entire neighborhoods and killed more than 13,000 people, according to heath authorities in Gaza. More than half of the territory’s 2.2 million Palestinians have been displaced, and the destruction means that many will have no homes to return to after the war ends.

On Wednesday, President Biden appeared to couch his otherwise strong embrace of Israel by suggesting that more fighting would benefit Hamas.

“Hamas unleashed a terrorist attack because they fear nothing more than Israelis and Palestinians living side by side in peace,” Mr. Biden said in a post on X, formerly Twitter. “To continue down the path of terror, violence, killing and war is to give Hamas what they seek. We can’t do that.”

Patrick Kingsley contributed reporting from Jerusalem, and Michael Crowley from Brussels.

Hamas released 10 Israeli hostages and two Thai nationals on Tuesday as the cease-fire between Israel and Hamas continued into a fifth day, raising the total number of captives released to 85.Roughly 240 hostages were captured on Oct. 7, when Hamas led devastating raids into Israel and killed about 1,200 people, according to the Israeli authorities.

Tuesday’s release followed the return of 11 Israelis on Monday, 13 on Sunday, 13 on Saturday and 13 on Friday, nearly all of them women or children. Seventeen Thais, one Filipino and one Russian-Israeli dual citizen were also released through separate negotiations. Four other hostages were released before the cease-fire.

Israel has freed 180 imprisoned Palestinians — roughly three for every Israeli hostage released — in exchange for the release of the hostages.

At least four hostages are believed to have died in captivity, according to Israel and Hamas. Israeli officials said that three — a 19-year-old soldier, Noa Marciano; Yehudit Weiss, 65; and Ravid Katz, 51 — were killed by their captors. Arye Zalmanovich, 86, died after suffering a heart attack, Hamas has said.

Israel and Hamas agreed on Monday to extend their truce to six days — through at least Wednesday — to allow for the release of more Israeli hostages and imprisoned Palestinians, and for more humanitarian aid to flow into the Gaza Strip, where thousands of Gazans have been killed and conditions for the living have deteriorated to the point of catastrophe.

Since the cease-fire began, Israel has continued its escalation of nighttime raids in East Jerusalem and the occupied West Bank, arresting at least 71 more Palestinians since Nov. 24, according to the Israeli military. The Palestinian Authority’s commission for prisoner affairs put the number of those arrested at 112 or more, including women, children and older people.

Image

Avigail Idan, in a photo provided by her family.
An aunt of Avigail Idan, the girl who was taken hostage by Hamas after she saw her parents brutally killed and who turned 4 a few days before being released, says that her niece shared one piece of pita bread per day with four others, and did not have a shower or bath during her 50 days in captivity.

Avigail, who is a dual citizen of Israel and the United States, was kept with the four members of the Brodutch family with whom she was kidnapped on Oct. 7. The Brodutches were the Idan family’s neighbors at the Kfar Aza kibbutz, and Avigail had hidden with the family after her parents were shot.

According to the aunt, Tal Idan, the five hostages were kept in aboveground apartments, changing locations at least once. They were given a piece of pita with za’atar, a Middle Eastern spice mixture, each day to share.

While in captivity, Avigail’s hair was cut because she had developed a significant case of lice, Ms. Idan said in an interview this week. “She was covered in it. It took quite an effort to help her get rid of some of it the first night,” she said.

“It was lovely curly hair. And now it’s all gone. But it will grow again.”

When Avigail was released on Sunday, she was met at the border in Israel by her grandmother and an aunt. A military helicopter took them to a pediatric hospital near Tel Aviv. The military gave her a set of small, pink noise-canceling headphones to put over her ears.

Avigail’s family was relieved to learn that she was not alone during her captivity, and was looked after by Hagar Brodutch as she also cared for her own three children. “We are so grateful for Hagar,,” Ms. Idan said, calling the 40-year-old “our guardian angel.”

Since the Oct. 7 attacks, Ms. Idan and her husband have helped care for Avigail’s siblings, Michael, 9, and Amelia, 6, who were not kidnapped. Now, she said, the children must adjust to a life without their parents.

“For all three of them, the focus is helping them get a new life — being able to feel safe again and being able to sleep for the entire night and not having nightmares,” she said.

Video

People are stocking up on basic necessities like water and fuel during the temporary window of relief from the violence.CreditCredit...Samar Abu Elouf for The New York Times
Gazans have been out on the streets during the temporary cease-fire between Israel and Hamas to stock up on supplies. Long lines snaked out of gas stations in Khan Younis this week, while merchants in Gaza City set up food stalls amid piles of rubble. But residents say the wait is long and many goods are in short supply.

Image

Israel’s national security minister, Itamar Ben-Gvir, center, in Jerusalem earlier this month.Credit...Ammar Awad/Reuters
As international pressure grows to extend a temporary cease-fire with Hamas, some right-wing members of Prime Minister Benjamin Netanyahu’s government are threatening to bring it down if he does not resume fighting in Gaza.

The far-right national security minister, Itamar Ben-Gvir, said on Wednesday that if Israel did not continue its war with Hamas, his political faction would leave the government coalition.

“Stopping the war = breaking apart the government,” said Mr. Ben-Gvir in a written statement.

While Mr. Ben-Gvir’s departure alone would not topple the government, it would give Mr. Netanyahu a very slim majority to keep his hold on power.

On Monday, Israel and Hamas agreed to prolong their cease-fire to six days from four, under a deal that would see the ongoing exchange of Palestinian prisoners for Israeli hostages, and mediators are trying to hammer out another extension. The longer pause was largely cheered by the Israeli public, which has been watching round-the-clock news coverage documenting the return of Israeli civilians who were kidnapped from their homes on Oct. 7.

But far-right members of Mr. Netanyahu’s government have been critical of the cease-fire, arguing that Israel should continue its military assault on Gaza. Mr. Ben-Gvir, who went from being a fringe figure in the Israeli settler movement to his current role in Mr. Netanyahu’s government, has been especially vocal, at times calling for Israel to “eliminate” anyone who supports Hamas.

Mr. Netanyahu can ill afford to alienate Mr. Ben-Gvir, who is part of a right-wing coalition of parties who give the prime minister a slim majority in the parliament. If he fails to hold a majority of the 120 seats, he would need to try to form a new coalition, or face another national election.

Two members of Mr. Netanyahu’s staff, who spoke on condition of anonymity because they were not authorized to speak to reporters, said the Israeli prime minister wanted to avoid elections at any cost.

Mr. Netanyahu’s approval ratings have steadily declined since the Oct. 7 attacks on Israel. In a poll conducted by Israel’s Bar-Ilan University earlier this month, trust in Mr. Netanyahu was at 4 percent. In another poll, conducted by the Israeli Maariv newspaper last week, 57 percent of Israelis said they would vote for Benny Gantz, a moderate member of Mr. Netanyahu’s war cabinet, for prime minister, over the 27 percent who said they would vote for Mr. Netanyahu.

Mr. Netanyahu has also been conducting his own polling, said the two staff members, and felt he would not fare well in a national election if it were held in the coming months.

— Sheera Frenkel reporting from Tel Aviv

Image

In this photo provided by the government of Israel, Natalie Raanan, 17, left, and her mother, Judith Raanan, 59, right, are being returned to Israel on Oct. 20.Credit...Government of Israel, via Associated Press
Hostages freed by Hamas are experiencing bursts of euphoria followed by emotional crashes as they emerge from weeks of paralyzing fear in captivity, according to Israeli relatives who spoke to journalists at the Israeli Embassy in London on Tuesday.

“Everything is just so fragile,” said Limor Sella-Broyde, 40, whose cousins Judith Raanan, 59, and Natalie Raanan, 17, were freed on Oct. 20.

She said the mother and daughter were in difficult psychological states as they recovered. “It’s an adrenaline rush and then a crash,” Ms. Sella-Broyde said. “They try to speak and then they get really tired.”

Three of Ms. Sella-Broyde’s relatives were killed by Hamas during the group’s terrorist attack in Israel on Oct. 7, while seven others were taken hostage from Kibbutz Be’eri.

Judith and Natalie Raanan, who are dual citizens of Israel and the United States, were the first hostages to be freed from captivity after negotiations involving the United States, Qatar and other nations. But for their family, any sense of relief remains out of reach because so many other members remain captive, and because the life they knew in the kibbutz has been shattered.

Ms. Sella-Broyde was part of a small group of Israelis brought to London this week to speak about their experiences related to the Oct. 7 attacks and to raise the profiles of the many other hostages who remain in captivity.

The event, sponsored by an Israeli charity, appeared designed in part to refocus public attention on Oct. 7, at a time when much of that attention has shifted to Israeli airstrikes and military operations in Gaza that have caused thousands of civilian deaths.

But the Israelis who spoke on Tuesday were reluctant to give details of the ordeals their family members had gone through in Gaza or describe their medical conditions after being released, citing a desire to protect their privacy.

Ms. Sella-Broyde said that she was nervous about broaching difficult topics with Judith and Natalie Raanan. “They are broken,” she said.

Another member of the group who spoke, Shahar Mor, 52, said that four of his family members had been kidnapped on Oct. 7. Three of them, including Ohad Munder Zichri, 9, have since been released, but the family’s patriarch, Avraham Munder, 78, has not.

Mr. Mor said that when his relatives were first freed, the family members simply hugged one another. But he added that he was dogged by a feeling of horror at the events of Oct. 7 and by worry for Mr. Munder.

In captivity, Mr. Mor said, his relatives had subsisted mainly on pita bread, which became more scarce as the weeks went by, and they ate with their guards. One guard gave them a pack of cards, which helped to pass the time, he said.

Mr. Mor downplayed the significance of the gesture, arguing that one humane act by a Hamas guard should not be a distraction from the group’s ruthlessness.

“Some of the captors are nicer; some are less nice,” he said. “It makes no difference.”

Image

Israeli soldiers heading south near Ashkelon, Israel, on Oct. 7.Credit...Ohad Zwigenberg/Associated Press
Israel confirmed on Tuesday the deaths of three Israeli soldiers who fell in combat during the Hamas-led terrorist attacks on Israel on Oct. 7, and a spokesman for the Israel Defense Forces press office, Nir Dinar, said Hamas was holding some of their remains.

The Israeli army identified the three soldiers as Sgt. Shaked Dahan, 19, of Afula, in northern Israel; Sgt. Kiril Brodski, 19, of Ramat Gan, near Tel Aviv; and Staff Sgt. Tomer Yaakov Ahimas, 20, from Lehavim, a town in southern Israel. The military ranks of all three soldiers had been upgraded after their deaths.

The army’s chief rabbi made a determination the soldiers had been killed during the Hamas-led attacks in which 1,200 people were massacred, based on evidence and intelligence reports, according to ynet.com, an Israeli news site.

Military funerals will be held on Wednesday for two of the soldiers, Sergeant Ahimas and Sergeant Brodski, ynet.com reported, suggesting enough remains were recovered and identified for funerals to be held for them under Jewish tradition.

Sigalit Gal, the mother of Sergeant Dahan, said in a Facebook post that she would not start the weeklong Jewish mourning ritual of Shiva for her son until his body was returned. “I fought all my life to take care and raise you with love, excellent education, values, a proper environment,” she wrote. “You were taken from me forever. They took you and did not care to return you — not even your body.”

It was not immediately clear why the army press office made the announcement on Tuesday, or whether the announcement might indicate that the soldiers’ bodies would be linked to the release of further Israeli hostages from Gaza.

Israel has in the past been willing to pay dearly in order to return bodies of its fallen soldiers to their families for burial, often releasing large numbers of Palestinian prisoners in exchange for bodies. In 2008, for example, Israel released five prisoners, including a Lebanese man convicted of murder, in order to obtain the bodies of two fallen Israeli soldiers.

Image

Israeli soldiers doing repairs near the border with the Gaza Strip on Tuesday.Credit...Christophe Petit Tesson/EPA, via Shutterstock
Under pressure from allies to ease the humanitarian crisis in Gaza as it welcomes Israeli hostages home under the terms of a fragile cease-fire, Israel faces an increasingly difficult set of decisions about the future of its war against Hamas.

Israeli leaders have vowed to eliminate Hamas, the group that has controlled Gaza since 2007 and that led devastating attacks on Israel on Oct. 7. They have also promised to recover all of the roughly 240 people who were kidnapped by Hamas and other Palestinian armed groups that day.

Prime Minister Benjamin Netanyahu has cited the recovery of hostages in justifying his support for the pause in Israel’s ground invasion — and has also said that Israel’s military is ready to resume fighting once the cease-fire arrangement ends.

But the deal also gives Hamas time to regroup and retrench, making Israel’s objective of rooting it out more difficult. And Israel’s release of Palestinians from imprisonment or detention under the arrangement has been accompanied by growing support for Hamas in the Israeli-occupied West Bank.

The extended cease-fire has also allowed aid to reach more of Gaza’s 2.2 million residents, most of whom have been displaced by the fighting and face dire shortages of food, medicine and fuel.

Where does the cease-fire agreement stand?
Israel and Hamas have extended their brief truce from four to six days, according to Qatar, which has been mediating their talks. The agreement has so far held, despite accusations by each side that the other had violated it.

Since Friday, Hamas has released at least 60 Israeli hostages and Israel has freed 180 imprisoned Palestinians. Nineteen other hostages in Gaza — 17 Thais, one Filipino and one Russian-Israeli dual citizen — have been released since Friday through separate negotiations.

Where does Israel’s military campaign stand?
Before the cease-fire agreement took hold on Friday, Israel’s military had bombarded the Gaza Strip for weeks, saying it had struck over 15,000 targets. The bombardment included the use of very large weapons in dense urban areas, and Palestinian health authorities in the Hamas government have said that more than 13,000 people have been killed, including thousands of children. The bombardment, the large number of deaths and the displacement of the majority of Gaza’s 2.2 million people have fueled international outcry over the scope of Israel’s campaign.

Israel has said it is targeting Hamas all over Gaza, including in places where its members are embedded among civilians, like hospitals and shelters, and in an extensive tunnel network underground.

Early in the war, the Israeli military ordered the evacuation of northern Gaza, ahead of a ground invasion. More than a million people fled south, and the invasion began in late October.

Israeli troops have captured a swath of northern Gaza roughly in the shape of a C: the northern edge of the strip, a sliver along the Mediterranean coast, and the central strip below Gaza City. The forces largely encircled Gaza City and split the strip in two halves, seeking to disrupt Hamas’s grip over the enclave and begin ousting it from its biggest city. But there appear to be parts of northern Gaza that the Israeli military does not control.

Israeli forces have also closed in on Gaza’s hospitals, seizing Al-Shifa, Gaza’s largest and most modern. Israel said Hamas used tunnels beneath the hospital as a command center, accusations the group and hospital staff denied.

The Israeli military has since sought to show evidence of its assertions, releasing videos that show parts of a tunnel shaft on the grounds of the Shifa complex, and rooms within the tunnel. But Israeli troops have moved slowly, wary of explosives and traps.

The military has said it has destroyed some Hamas tunnels, but has not said that its troops have been fighting inside them.

Ron Dermer, a member of Israel’s war cabinet, told Sky News on Nov. 7 that the Israeli military had killed “several thousand” Hamas fighters since the war began. He said that the total at that point was greater than 3,000 and “probably close to 4,000 already.” Israeli officials estimate that Hamas had numbered about 25,000 members before the war began.

More than 70 Israeli soldiers have been killed since the ground invasion began, according to the Israeli military.

What about Hamas’s leadership?
Hamas has acknowledged the deaths of several commanders in the war, including at least one senior figure. A number of other Hamas officials and commanders are believed to have been killed. Yahya Sinwar, the hard-line leader of Hamas in Gaza, remains a top target of Israeli forces.

Israeli leaders have said they do not want to reoccupy Gaza after the war, and it remains unclear how or whether they can eliminate Hamas completely from the strip. And in the West Bank, which the Israeli military occupies, support for Hamas has grown amid the recent releases of Palestinians and mounting frustration with the Palestinian Authority, which oversees the West Bank.

The political leadership of Hamas is not within Israel’s reach. Qatar hosts Hamas’s political leaders in its capital, Doha, where Qatari officials have been mediating the talks between Israel and Hamas alongside Egypt and the United States.`,
		},
	}
	input_stream := spider.GetInputStream()
	for _, topic := range topics {
		item := topic.Content
		log.Infof("Get Item: %s ...", item[:20])
		input_stream.Write(topic)
	}
}

func TestChatGPT(t *testing.T) {
	var conf gaddets.Config
	err := util.ParseConfig("../config.yml", &conf)
	if err != nil {
		log.Panic(err)
	}
	//sixtone := spider.NewSixtoneSpider()
	newsSpider := NewNewsSpider()
	chatgpt := NewChatGPT(&conf)
	newsSpider.Chain(&chatgpt)
	Exec.Add(&newsSpider)
	Exec.Start()
}
