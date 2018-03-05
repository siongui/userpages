[Golang] Parse Dash Manifest in Instagram Post Live Story
#########################################################

:date: 2018-03-05T22:36+08:00
:tags: Go, Golang, Regular Expression, String Manipulation, Instagram, XML, html
:category: Go
:summary: Get urls of video and audio of Instagram user post live videos shared
          to stories via Go standard *regexp* package.
:og_image: http://slideplayer.com/slide/4557119/15/images/17/DASH+Manifest+(MPD)+high-level+structure+(on-demand+profile).jpg
:adsu: yes


Get urls of video and audio of Instagram_ user post live videos that are shared
to stories. We will parse `dash manifest`_ via Go standard regexp_ package.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ZwrYqxc_l7X>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"regexp"
  )

  const testDashManifest = `<MPD xmlns="urn:mpeg:dash:schema:mpd:2011" minBufferTime="PT1.500S" type="static" mediaPresentationDuration="PT0H0M9.468S" maxSegmentDuration="PT0H0M2.000S" profiles="urn:mpeg:dash:profile:isoff-on-demand:2011,http://dashif.org/guidelines/dash264"><Period duration="PT0H0M9.468S"><AdaptationSet segmentAlignment="true" maxWidth="396" maxHeight="746" maxFrameRate="16000/528" par="396:746" lang="und" subsegmentAlignment="true" subsegmentStartsWithSAP="1"><Representation id="17924879014006631v" mimeType="video/mp4" codecs="avc1.4d401f" width="396" height="746" frameRate="16000/528" sar="1:1" startWithSAP="1" bandwidth="836675" FBQualityClass="sd" FBQualityLabel="396w"><BaseURL>https://instagram.fkhh1-1.fna.fbcdn.net/vp/4bc800aa48ed9f3bcc763d3e5d4e48fe/5A8F0D5A/t72.12950-16/27465973_336693050159819_3428588455850934272_n.mp4</BaseURL><SegmentBase indexRangeExact="true" indexRange="899-1026"><Initialization range="0-898"/></SegmentBase></Representation></AdaptationSet><AdaptationSet segmentAlignment="true" lang="und" subsegmentAlignment="true" subsegmentStartsWithSAP="1"><Representation id="17924879014006631a" mimeType="audio/mp4" codecs="mp4a.40.2" audioSamplingRate="44100" startWithSAP="1" bandwidth="51679"><AudioChannelConfiguration schemeIdUri="urn:mpeg:dash:23003:3:audio_channel_configuration:2011" value="2"/><BaseURL>https://instagram.fkhh1-1.fna.fbcdn.net/vp/c3198573a9e6375f6bd01da3c8312dc4/5A8F0753/t72.12950-16/27486167_191963884733826_1789728450089582592_n.mp4</BaseURL><SegmentBase indexRangeExact="true" indexRange="835-926"><Initialization range="0-834"/></SegmentBase></Representation></AdaptationSet></Period></MPD>`

  func getBaseUrls(dm string) (urls []string, err error) {
  	reBaseUrls, err := regexp.Compile(`<BaseURL>(.+?)<\/BaseURL>`)
  	if err != nil {
  		return
  	}

  	matches := reBaseUrls.FindAllStringSubmatch(dm, -1)
  	for _, match := range matches {
  		urls = append(urls, match[1])
  	}
  	return
  }

  func main() {
  	urls, err := getBaseUrls(testDashManifest)
  	if err != nil {
  		panic(err)
  	}

  	for _, url := range urls {
  		// 2 urls: one is video, and the other is audio
  		fmt.Println(url)
  	}
  }

.. adsu:: 2

Tested on: ``Ubuntu 17.10 + Go 1.10`` and `Go Playground`_

----

References:

.. [1] `Online regex tester and debugger: PHP, PCRE, Python, Golang and JavaScript <https://regex101.com/>`_
.. [2] | `regex non greedy - Google search <https://www.google.com/search?q=regex+non+greedy>`_
       | `regex non greedy - DuckDuckGo search <https://duckduckgo.com/?q=regex+non+greedy>`_
       | `regex non greedy - Ecosia search <https://www.ecosia.org/search?q=regex+non+greedy>`_
       | `regex non greedy - Qwant search <https://www.qwant.com/?q=regex+non+greedy>`_
       | `regex non greedy - Bing search <https://www.bing.com/search?q=regex+non+greedy>`_
       | `regex non greedy - Yahoo search <https://search.yahoo.com/search?p=regex+non+greedy>`_
       | `regex non greedy - Baidu search <https://www.baidu.com/s?wd=regex+non+greedy>`_
       | `regex non greedy - Yandex search <https://www.yandex.com/search/?text=regex+non+greedy>`_
       |
       | `How can I write a regex which matches non greedy? - Stack Overflow <https://stackoverflow.com/a/11899069>`_
.. [3] | `golang regex group example - Google search <https://www.google.com/search?q=golang+regex+group+example>`_
       | `golang regex group example - DuckDuckGo search <https://duckduckgo.com/?q=golang+regex+group+example>`_
       | `golang regex group example - Ecosia search <https://www.ecosia.org/search?q=golang+regex+group+example>`_
       | `golang regex group example - Qwant search <https://www.qwant.com/?q=golang+regex+group+example>`_
       | `golang regex group example - Bing search <https://www.bing.com/search?q=golang+regex+group+example>`_
       | `golang regex group example - Yahoo search <https://search.yahoo.com/search?p=golang+regex+group+example>`_
       | `golang regex group example - Baidu search <https://www.baidu.com/s?wd=golang+regex+group+example>`_
       | `golang regex group example - Yandex search <https://www.yandex.com/search/?text=golang+regex+group+example>`_
       |
       | `Golang-Regex-Tutorial/01-chapter2.markdown at master · StefanSchroeder/Golang-Regex-Tutorial · GitHub <https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter2.markdown>`_
       | `func (*Regexp) FindAllStringSubmatch - regexp - The Go Programming Language <https://golang.org/pkg/regexp/#Regexp.FindAllStringSubmatch>`_
.. [4] `get url of video and audio of post live · siongui/instago@eb122f5 · GitHub <https://github.com/siongui/instago/commit/eb122f557a697c721538d9e88090cc3dbda09538>`_

.. _regexp: https://golang.org/pkg/regexp/
.. _Instagram: https://www.instagram.com/
.. _dash manifest: https://www.google.com/search?q=dash+manifest
.. _Go Playground: https://play.golang.org/
