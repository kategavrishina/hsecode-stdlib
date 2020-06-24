package fulltext_test

import (
	fulltext "hsecode.com/stdlib/strings/fulltext"
	"testing"
)

func TestFullText(t *testing.T) {
	docs := fulltext.New([]string{
		"this is the house that jack built",
		"this the rat that ate the malt is",
		"this the is the",
		"rate the eat this jack is",
		"rate the eat this jack is",
	})
	t.Log(docs)
	s0 := docs.Search("")
	if len(s0) > 0 {
		t.Fatalf("Empty search doesn't work: %v", s0)
	}
	s1 := docs.Search("in the house that jack built")
	if len(s1) > 0 {
		t.Fatalf("No documents found doesn't work: %v", s1)
	}
	s2 := docs.Search("malt rat")
	if len(s2) != 1 {
		t.Fatalf("One document found doesn't work: %v", s2)
	}
	s3 := docs.Search("is this the")
	if s3[1] != 1 {
		t.Fatalf("Two documents found doesn't work: %v", s3)
	}
	if len(s3) != 5 {
		t.Fatalf("Two documents found doesn't work: %v", s3)
	}
	t.Log(s3)
}

func TestDuplicate(t *testing.T) {
	docs := fulltext.New([]string{
		"this is the house that jack built",
		"house that jack built this is the",
	})
	s0 := docs.Search("the jack")
	if s0[1] != 1 {
		t.Fatalf("Duplicate documents doesn't work: %v", s0)
	}
	if len(s0) != 2 {
		t.Fatalf("Duplicate documents doesn't work: %v", s0)
	}
}

func TestLarge(t *testing.T) {
	docs := fulltext.New([]string{
		"fghj rtyui cvb ertyui fghjk fghj rtyui this is the house that jack built cvb ertyui fghjkfghj rtyui cvb ertyui this is the house that jack built fghjk fghj rtyui cvb ertyui fghjk",
		"this is the house that jack built wert vbnm, iop[ sdfg xcvbnm tyui ;lkjh tre bvc fds oiuy fds mnbv fds ouyf ikhbvrb iythgnbv",
		"rate the eat this jack is kxejf okjgrv kj vksjfa pakfn palej paoeri lxkjhf paoitru kjfhg paiu lkf paoei kz dutr",
		"Algorithm In the previous post, we discussed construction of BST from sorted Linked List. Constructing from sorted array in O(n) time is simpler as we can get the middle element in O(1) time. Following is a simple algorithm where we first find the middle node of list and make it root of the tree to be constructed. 1) Get the Middle of the array and make it root. 2) Recursively do same for left half and right half. a) Get the middle of left half and make it left child of the root created in step 1. b) Get the middle of right half and make it right child of the root created in step 1. Following is the implementation of the above algorithm. The main code which creates Balanced BST is highlighted.",
		"ou ouer zmdfh seg nxhdf kz sdfknv kjzhdf nbbuses kzjyg kzjhgf zkjhdg zkjsfg owwery zkjdhg z,jkdht iauygr jhzdfvb kshvfhds",
		"In this paper we presented a novel architecture for NER that expands the feature set space based on feature clustering of images and texts, focused on microblogs. Due to their terse nature, such noisy data often lack enough context, which poses a challenge to the correct identification of named entities. To address this issue we have presented and evaluated a novel approach using the Ritter dataset, showing consistent results over state-of-the-art models without using any external resource or encoded rule, achieving an average of 0.59 F1. The results slightly outperformed state-of-the-art models which do not rely on encoded rules (0.49 and 0.54 F1), suggesting the viability of using the produced metadata to also boost existing NER approaches. A further important contribution is the ability to handle single tokens and misspelled words successfully, which is of utmost importance in order to better understand short texts. Finally, the architecture of the approach and its indicators introduce potential to transparently support multilingual data, which is the subject of ongoing investigation.",
		"In this paper we presented a novel architecture for NER that expands the feature set space based on feature clustering of images and texts, focused on microblogs. Due to their terse nature, such noisy data often lack enough context, which poses a challenge to the correct identification of named entities. To address this issue we have presented and evaluated a novel approach using the Ritter dataset, showing consistent results over state-of-the-art models without using any external resource or encoded rule, achieving an average of 0.59 F1. The results slightly outperformed state-of-the-art models which do not rely on encoded rules (0.49 and 0.54 F1), suggesting the viability of using the produced metadata to also boost existing NER approaches. A further important contribution is the ability to handle single tokens and misspelled words successfully, which is of utmost importance in order to better understand short texts. Finally, the architecture of the approach and its indicators introduce potential to transparently support multilingual data, which is the subject of ongoing investigation.",
	})
	t.Log(docs.Search("is"))
}
