// Code generated by "stringer -type=DeviceType,BrowserName,OSName,Platform -output=const_string.go"; DO NOT EDIT.

package uasurfer

import "fmt"

const _DeviceType_name = "DeviceUnknownDeviceComputerDeviceTabletDevicePhoneDeviceConsoleDeviceWearableDeviceTV"

var _DeviceType_index = [...]uint8{0, 13, 27, 39, 50, 63, 77, 85}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceType_index)-1) {
		return fmt.Sprintf("DeviceType(%d)", i)
	}
	return _DeviceType_name[_DeviceType_index[i]:_DeviceType_index[i+1]]
}

const _BrowserName_name = "BrowserUnknownBrowserChromeBrowserIEBrowserSafariBrowserFirefoxBrowserAndroidBrowserOperaBrowserBlackberryBrowserUCBrowserBrowserSilkBrowserNokiaBrowserNetFrontBrowserQQBrowserMaxthonBrowserSogouExplorerBrowserSpotifyBrowserNintendoBrowserSamsungBrowserYandexBrowserCocCocBrowserBotBrowserAppleBotBrowserBaiduBotBrowserBingBotBrowserDuckDuckGoBotBrowserFacebookBotBrowserGoogleBotBrowserLinkedInBotBrowserMsnBotBrowserPingdomBotBrowserTwitterBotBrowserYandexBotBrowserCocCocBotBrowserYahooBot"

var _BrowserName_index = [...]uint16{0, 14, 27, 36, 49, 63, 77, 89, 106, 122, 133, 145, 160, 169, 183, 203, 217, 232, 246, 259, 272, 282, 297, 312, 326, 346, 364, 380, 398, 411, 428, 445, 461, 477, 492}

func (i BrowserName) String() string {
	if i < 0 || i >= BrowserName(len(_BrowserName_index)-1) {
		return fmt.Sprintf("BrowserName(%d)", i)
	}
	return _BrowserName_name[_BrowserName_index[i]:_BrowserName_index[i+1]]
}

const _OSName_name = "OSUnknownOSWindowsPhoneOSWindowsOSMacOSXOSiOSOSAndroidOSBlackberryOSChromeOSOSKindleOSWebOSOSLinuxOSPlaystationOSXboxOSNintendoOSBot"

var _OSName_index = [...]uint8{0, 9, 23, 32, 40, 45, 54, 66, 76, 84, 91, 98, 111, 117, 127, 132}

func (i OSName) String() string {
	if i < 0 || i >= OSName(len(_OSName_index)-1) {
		return fmt.Sprintf("OSName(%d)", i)
	}
	return _OSName_name[_OSName_index[i]:_OSName_index[i+1]]
}

const _Platform_name = "PlatformUnknownPlatformWindowsPlatformMacPlatformLinuxPlatformiPadPlatformiPhonePlatformiPodPlatformBlackberryPlatformWindowsPhonePlatformPlaystationPlatformXboxPlatformNintendoPlatformBot"

var _Platform_index = [...]uint8{0, 15, 30, 41, 54, 66, 80, 92, 110, 130, 149, 161, 177, 188}

func (i Platform) String() string {
	if i < 0 || i >= Platform(len(_Platform_index)-1) {
		return fmt.Sprintf("Platform(%d)", i)
	}
	return _Platform_name[_Platform_index[i]:_Platform_index[i+1]]
}